package appsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

func NewApi(mod module.Module, alias string) awsappsync.CfnGraphQLApi {

	name := naming.NewName(mod, naming.ResType_GraphQLApi, alias)

	return awsappsync.NewCfnGraphQLApi(
		mod.Resource(),
		name.Logical(),
		&awsappsync.CfnGraphQLApiProps{
			Name:               name.Logical(),
			AuthenticationType: jsii.String("API_KEY"),
			XrayEnabled:        jsii.Bool(true),
		},
	)
}
