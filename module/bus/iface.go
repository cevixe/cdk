package bus

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/cevixe/cdk/module"
)

type BusType uint8

const (
	BusType_Standard BusType = 0
	BusType_Advanced BusType = 1
)

type Bus interface {
	Module() module.Module
	Name() string
	Type() BusType
	Resource() awssns.Topic
}
