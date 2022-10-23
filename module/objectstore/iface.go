package objectstore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/cevixe/cdk/module"
)

type ObjectStore interface {
	Module() module.Module
	Name() string
	Resource() awss3.Bucket
}
