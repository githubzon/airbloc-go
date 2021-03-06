package storage

import (
	"net/url"

	"github.com/airbloc/airbloc-go/data"
)

type Type int8

const (
	Local Type = iota
	CloudS3
	CloudGoogle
	CloudAzure
)

var Type_value = map[string]Type{
	"local": Local,
	"s3":    CloudS3,
}

var Type_name = map[Type]string{
	Local:   "local",
	CloudS3: "s3",
}

type Storage interface {
	Save(string, *data.Bundle) (*url.URL, error)
	Update(*url.URL, *data.Bundle) error
	Delete(*url.URL) error
}
