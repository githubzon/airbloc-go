package p2p

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/azer/logger"

	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/p2p/common"
	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	"github.com/gogo/protobuf/proto"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p/p2p/host/routed"
	"github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"
	"github.com/pkg/errors"
)

const (
	ProtocolName    = "airbloc"
	ProtocolVersion = "1.0.0"
)

type AirblocServer struct {
	// controller
	mutex  *sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc

	// network
	id      cid.Cid
	pid     common.Pid
	host    Host
	dht     *kaddht.IpfsDHT
	nodekey *key.Key

	// topic - handlers
	types    map[string]reflect.Type
	handlers map[string]TopicHandler

	// log
	log *logger.Logger
}

func NewAirblocServer(
	nodekey *key.Key,
	addr multiaddr.Multiaddr,
	bootinfos []peerstore.PeerInfo,
) (Server, error) {
	if len(bootinfos) < 1 {
		return nil, errors.New("bootnode informations should be give at least one.")
	}

	privKey, err := nodekey.DeriveLibp2pKeyPair()
	if err != nil {
		return nil, err
	}

	pid, err := common.NewPid(ProtocolName, ProtocolVersion)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate pid")
	}

	ctx, cancel := context.WithCancel(context.Background())
	server := &AirblocServer{
		ctx:     ctx,
		cancel:  cancel,
		mutex:   new(sync.Mutex),
		pid:     pid,
		nodekey: nodekey,

		types:    make(map[string]reflect.Type),
		handlers: make(map[string]TopicHandler),
		log:      logger.New("p2p"),
	}

	h, err := libp2p.New(
		ctx,
		libp2p.Identity(privKey),
		libp2p.ListenAddrs(addr),
	)
	if err != nil {
		cancel()
		return nil, err
	}

	server.dht, err = kaddht.New(ctx, h)
	if err != nil {
		cancel()
		return nil, err
	}

	h = routedhost.Wrap(h, server.dht)
	server.host = NewAirblocHost(NewBasicHost(h), 20)

	// connect to bootstrap nodes for initializing DHT
	for _, bootinfo := range bootinfos {
		if err := h.Connect(ctx, bootinfo); err != nil {
			cancel()
			return nil, errors.Wrap(err, "failed to connect to bootstrap node")
		}
	}

	idVal := int32(pb.CID_AIRBLOC)

	v1b := cid.V1Builder{
		Codec:  uint64(idVal),
		MhType: multihash.KECCAK_256,
	}

	server.id, err = v1b.Sum([]byte(pb.CID_name[idVal]))
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "server error : failed to generate cid")
	}

	server.log.Info("Initialized", logger.Attrs{
		"protocol":   fmt.Sprintf("%s %s", ProtocolName, ProtocolVersion),
		"on address": addr.String(),
	})
	return server, nil
}

// Discovery finds and updates new peer connection every minute.
func (s *AirblocServer) Discovery() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	numOfPeers := 0
	s.updatePeer()
	for {
		select {
		case <-ticker.C:
			found := s.updatePeer()
			if numOfPeers != found {
				s.log.Info("Connected", logger.Attrs{"peers": found})
				numOfPeers = found
			}
		case <-s.ctx.Done():
			return
		}
	}
}

func (s *AirblocServer) clearPeer() {
	peerStore := s.host.Peerstore()
	for _, peerID := range peerStore.PeersWithAddrs() {
		peerStore.ClearAddrs(peerID)
	}
}

func (s *AirblocServer) updatePeer() int {
	idch, err := s.dht.GetClosestPeers(s.ctx, s.id.KeyString())
	if s.ctx.Err() != nil {
		s.log.Error("Failed to discovery peers: context error: %v", s.ctx.Err())
		return 0
	}

	if err != nil {
		s.log.Error("Failed to discovery peers: %v", err)
	}

	found := 0
	for id := range idch {
		info, err := s.dht.FindPeer(s.ctx, id)
		if err != nil {
			s.log.Error("Warning: Peer found, but cannot connect", logger.Attrs{"to": id.Pretty()})
			continue
		}
		s.host.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.TempAddrTTL)
		found++
	}
	return found
}

// api backend interfaces
func (s *AirblocServer) Start() error {
	s.host.RegisterProtocol(s.pid, s.handleMessage)
	go s.Discovery()
	return nil
}

func (s *AirblocServer) handleMessage(message common.ProtoMessage) {
	topic := message.GetTopic()
	typ, ok := s.types[topic]
	if !ok {
		s.log.Error("Unknown topic: %s", message.GetTopic())
		return
	}
	msg, err := message.MakeMessage(s.ctx, typ)
	if err != nil {
		s.log.Error("Failed to make message: %v", err.Error())
		return
	}

	timer := s.log.Timer()
	handler := s.handlers[topic]
	handler(s, s.ctx, msg)
	timer.End("Received message", logger.Attrs{
		"from":  msg.SenderAddr.String(),
		"topic": topic,
	})
}

func (s *AirblocServer) Stop() {
	s.cancel()
}

func (s *AirblocServer) SubscribeTopic(topic string, msg proto.Message, handler TopicHandler) error {
	typ := common.MessageType(msg)

	s.mutex.Lock()
	s.types[topic] = typ
	s.handlers[topic] = handler
	s.mutex.Unlock()

	return nil
}

func (s *AirblocServer) UnsubscribeTopic(topic string) error {
	s.mutex.Lock()
	delete(s.types, topic)
	delete(s.handlers, topic)
	s.mutex.Unlock()

	return nil
}

func (s *AirblocServer) Send(ctx context.Context, msg proto.Message, topic string, p peer.ID) error {
	s.log.Info("Sending P2P message", logger.Attrs{
		"topic": topic,
		"id":    p.Pretty(),
	})
	payload, err := common.NewProtoMessage(msg, topic)
	if err != nil {
		return errors.Wrap(err, "send error")
	}
	return s.host.Send(ctx, *payload, p, s.pid)
}

func (s *AirblocServer) Publish(ctx context.Context, msg proto.Message, topic string) error {
	s.log.Info("Broadcasting P2P message", logger.Attrs{"topic": topic})
	payload, err := common.NewProtoMessage(msg, topic)
	if err != nil {
		return errors.Wrap(err, "publish error")
	}
	return s.host.Publish(ctx, *payload, s.pid)
}

// for test
func (s *AirblocServer) setContext(ctx context.Context) {
	s.ctx = ctx
}

func (s *AirblocServer) getHost() Host {
	return s.host
}
