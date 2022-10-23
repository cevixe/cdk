package apollo

import (
	"fmt"

	"github.com/cevixe/cdk/common/file"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/appsync"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
)

type EntitiesResolverProps struct {
	Api        awsappsync.CfnGraphQLApi `field:"required"`
	DataSource awsappsync.CfnDataSource `field:"required"`
	Table      awsdynamodb.Table        `field:"required"`
}

func NewEntitiesResolver(mod module.Module, alias string, props *EntitiesResolverProps) awsappsync.CfnResolver {

	requestLocation := "libraries/cevixe/cdk/apollo/templates/entities/request.vtl"
	request := fmt.Sprintf(file.GetFileContent(requestLocation), *props.Table.TableName())

	responseLocation := "libraries/cevixe/cdk/apollo/templates/entities/response.vtl"
	response := file.GetFileContent(responseLocation)

	return appsync.NewResolver(
		mod,
		alias,
		&appsync.ResolverProps{
			Api:        props.Api,
			DataSource: props.DataSource,
			Type:       "Query",
			Field:      "_entities",
			Request:    jsii.String(request),
			Response:   jsii.String(response),
		},
	)
}
