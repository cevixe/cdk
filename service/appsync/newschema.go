package appsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

type SchemaProps struct {
	Api        awsappsync.CfnGraphQLApi `field:"required"`
	Definition *string                  `field:"required"`
}

func NewSchema(mod module.Module, alias string, props *SchemaProps) awsappsync.CfnGraphQLSchema {

	name := naming.NewName(mod, naming.ResType_GraphQLSchema, alias)

	return awsappsync.NewCfnGraphQLSchema(
		mod.Resource(),
		name.Logical(),
		&awsappsync.CfnGraphQLSchemaProps{
			ApiId:      props.Api.AttrApiId(),
			Definition: props.Definition,
		},
	)
}
