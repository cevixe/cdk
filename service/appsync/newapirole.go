package appsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/iam"
)

func NewApiRole(mod module.Module, alias string) awsiam.Role {
	return iam.NewServiceRole(mod, alias, "appsync.amazonaws.com")
}
