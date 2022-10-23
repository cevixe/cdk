package statestore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/cevixe/cdk/module"
)

type stateStoreImpl struct {
	module   module.Module
	name     string
	resource awsdynamodb.Table
}

func (store *stateStoreImpl) Module() module.Module {
	return store.module
}

func (store *stateStoreImpl) Name() string {
	return store.name
}

func (store *stateStoreImpl) Resource() awsdynamodb.Table {
	return store.resource
}
