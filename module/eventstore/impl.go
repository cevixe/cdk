package eventstore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/cevixe/cdk/module"
)

type eventStoreImpl struct {
	module   module.Module
	name     string
	resource awsdynamodb.Table
}

func (store *eventStoreImpl) Module() module.Module {
	return store.module
}

func (store *eventStoreImpl) Name() string {
	return store.name
}

func (store *eventStoreImpl) Resource() awsdynamodb.Table {
	return store.resource
}
