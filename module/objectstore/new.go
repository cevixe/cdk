package objectstore

import (
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/s3"
)

func NewObjectStore(mod module.Module, alias string) ObjectStore {

	bucket := s3.NewBucket(mod, alias)

	return &objectStoreImpl{
		module:   mod,
		name:     alias,
		resource: bucket,
	}
}
