package appsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

type ResolverProps struct {
	Api        awsappsync.CfnGraphQLApi `field:"required"`
	DataSource awsappsync.CfnDataSource `field:"required"`
	Type       string                   `field:"required"`
	Field      string                   `field:"required"`
	Request    *string                  `field:"required"`
	Response   *string                  `field:"required"`
}

func NewResolver(mod module.Module, alias string, props *ResolverProps) awsappsync.CfnResolver {

	name := naming.NewName(mod, naming.ResType_GraphQLResolver, alias)

	return awsappsync.NewCfnResolver(
		mod.Resource(),
		name.Logical(),
		&awsappsync.CfnResolverProps{
			ApiId:                   props.Api.AttrApiId(),
			DataSourceName:          props.DataSource.AttrName(),
			TypeName:                jsii.String(props.Type),
			FieldName:               jsii.String(props.Field),
			RequestMappingTemplate:  props.Request,
			ResponseMappingTemplate: props.Response,
		},
	)
}
