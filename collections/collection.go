package collections

import (
	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/schemas"
	"time"
)

type Collection struct {
	Id common.ID

	// owner of the collection
	AppId common.ID

	IncentivePolicy IncentivePolicy

	// format of the data the app want to collect
	Schema    schemas.Schema
	CreatedAt time.Time

	// display-purpose data (Saved in MetaDB)
	DisplayedName string
	Category      Category
	Label         string
}

// Category is a (displayed) group of data collections defined by application.
type Category struct {
	Id            string
	DisplayedName string
}

// IncentivePolicy is a policy about sharing revenue of the data
// between each stakeholders at the data delivery pipeline.
type IncentivePolicy struct {
	DataProvider  float64
	DataProcessor float64
	DataRelayer   float64
	DataOwner     float64
}

func NewCollection(appId common.ID, schemaId common.ID, policy IncentivePolicy) *Collection {
	collection := &Collection{
		AppId:           appId,
		IncentivePolicy: policy,
	}
	collection.Schema.Id = schemaId
	return collection
}
