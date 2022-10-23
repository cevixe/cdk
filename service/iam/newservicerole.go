package iam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

func NewServiceRole(mod module.Module, alias string, service string) awsiam.Role {

	name := naming.NewName(mod, naming.ResType_IAMRole, alias)

	return awsiam.NewRole(
		mod.Resource(),
		name.Logical(),
		&awsiam.RoleProps{
			RoleName:  name.Physical(),
			AssumedBy: awsiam.NewServicePrincipal(jsii.String(service), nil),
		},
	)
}
