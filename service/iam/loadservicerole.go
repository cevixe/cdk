package iam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

func LoadServiceRole(mod module.Module, alias string, roleArn string) awsiam.IRole {

	name := naming.NewName(mod, naming.ResType_IAMRole, alias)

	return awsiam.Role_FromRoleArn(
		mod.Resource(),
		name.Logical(),
		jsii.String(roleArn),
		&awsiam.FromRoleArnOptions{
			Mutable: jsii.Bool(true),
		},
	)
}
