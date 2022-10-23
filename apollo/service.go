package apollo

import (
	"fmt"
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/common/file"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/appsync"
)

type ServiceResolverProps struct {
	Api        awsappsync.CfnGraphQLApi `field:"required"`
	DataSource awsappsync.CfnDataSource `field:"required"`
	Schema     *string                  `field:"required"`
}

func NewServiceResolver(mod module.Module, alias string, props *ServiceResolverProps) awsappsync.CfnResolver {

	escapedSchema := escapeSchema(props.Schema)

	requestLocation := "libraries/cevixe/cdk/apollo/templates/service/request.vtl"
	request := fmt.Sprintf(file.GetFileContent(requestLocation), *escapedSchema)

	responseLocation := "libraries/cevixe/cdk/apollo/templates/service/response.vtl"
	response := file.GetFileContent(responseLocation)

	return appsync.NewResolver(
		mod,
		alias,
		&appsync.ResolverProps{
			Api:        props.Api,
			DataSource: props.DataSource,
			Type:       "Query",
			Field:      "_service",
			Request:    jsii.String(request),
			Response:   jsii.String(response),
		},
	)
}

func escapeSchema(schema *string) *string {
	escapedSchema := *schema
	escapedSchema = strings.ReplaceAll(escapedSchema, `"`, `\"`)
	escapedSchema = strings.ReplaceAll(escapedSchema, "\n", `\n`)
	return &escapedSchema
}
