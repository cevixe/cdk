package appsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

type ApiKeyProps struct {
	Api awsappsync.CfnGraphQLApi `field:"required"`
}

func NewApiKey(mod module.Module, alias string, props *ApiKeyProps) awsappsync.CfnApiKey {

	name := naming.NewName(mod, naming.ResType_GraphQLSchema, alias)

	return awsappsync.NewCfnApiKey(
		mod.Resource(),
		name.Logical(),
		&awsappsync.CfnApiKeyProps{
			ApiId:    props.Api.AttrApiId(),
			ApiKeyId: name.Physical(),
		},
	)
}
