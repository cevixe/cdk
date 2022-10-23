package function

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/cevixe/cdk/module"
)

type functionImpl struct {
	module   module.Module
	name     string
	resource awslambda.Function
}

func (fn *functionImpl) Module() module.Module {
	return fn.module
}

func (fn *functionImpl) Name() string {
	return fn.name
}

func (fn *functionImpl) Resource() awslambda.Function {
	return fn.resource
}
