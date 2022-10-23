package bus

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/cevixe/cdk/module"
)

type busImpl struct {
	module   module.Module
	name     string
	typ      BusType
	resource awssns.Topic
}

func (b *busImpl) Module() module.Module {
	return b.module
}

func (b *busImpl) Name() string {
	return b.name
}

func (b *busImpl) Type() BusType {
	return b.typ
}

func (b *busImpl) Resource() awssns.Topic {
	return b.resource
}
