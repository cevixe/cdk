package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/iam"
)

func NewFunctionRole(mod module.Module, alias string) awsiam.Role {
	return iam.NewServiceRole(mod, alias, "lambda.amazonaws.com")
}
