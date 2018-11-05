package collections

import (
	"github.com/airbloc/airbloc-go/api"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/context"
)

type API struct {
	service *Service
}

func NewAPI(backend *api.AirblocBackend) (api.API, error) {
	service, err := NewService(backend.LocalDatabase, backend.Ethclient, nil, common.Address{})
	return &API{service}, err
}

func (api *API) Create(ctx context.Context, req *CreateCollectionRequest) (*CreateCollectionResponse, error) {
	hash, err := api.service.Register(ctx, &Collection{
		AppId:    common.HexToHash(req.AppId),
		SchemaId: common.HexToHash(req.SchemaId),
		Policy: &IncentivePolicy{
			DataProducer:  req.Policy.DataProducer,
			DataProcessor: req.Policy.DataProcessor,
			DataRelayer:   req.Policy.DataRelayer,
			DataSource:    req.Policy.DataSource,
		},
	})
	return &CreateCollectionResponse{
		CollectionId: hash.Hex(),
	}, err
}

// TODO after localdb integration
func (api *API) List(ctx context.Context, req *ListCollectionRequest) (*ListCollectionResponse, error) {
	return nil, nil
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterCollectionServer(service.GrpcServer, api)
}
