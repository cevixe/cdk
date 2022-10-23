package objectstore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/cevixe/cdk/module"
)

type objectStoreImpl struct {
	module   module.Module
	name     string
	resource awss3.Bucket
}

func (o *objectStoreImpl) Module() module.Module {
	return o.module
}

func (o *objectStoreImpl) Name() string {
	return o.name
}

func (o *objectStoreImpl) Resource() awss3.Bucket {
	return o.resource
}
