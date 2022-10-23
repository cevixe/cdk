package statestore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/cevixe/cdk/module"
)

type StateStore interface {
	Module() module.Module
	Name() string
	Resource() awsdynamodb.Table
}
