package p2p

import (
	"context"
	"log"

	"github.com/airbloc/airbloc-go/p2p/common"
	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
)

func testPingHandler(s Server, ctx context.Context, message common.Message) {
	log.Println("Ping", message.SenderInfo.ID.Pretty(), message.Data.String())

	s.Send(ctx, &pb.TestPing{Message: "World!"}, "ping", message.SenderInfo.ID)
}

func testPongHandler(s Server, ctx context.Context, message common.Message) {
	log.Println("Pong", message.SenderInfo.ID.Pretty(), message.Data.String())
}
