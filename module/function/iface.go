package function

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/cevixe/cdk/module"
)

type Function interface {
	Module() module.Module
	Name() string
	Resource() awslambda.Function
}
