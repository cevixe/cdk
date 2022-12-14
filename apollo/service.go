package apollo

import (
	"fmt"
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/appsync"
)

type ServiceResolverProps struct {
	Api        awsappsync.CfnGraphQLApi `field:"required"`
	DataSource awsappsync.CfnDataSource `field:"required"`
	Schema     *string                  `field:"required"`
}

const ServiceRequest = `
{
    "version": "2017-02-28",
	"payload": {
		"sdl": "%s"
	}
}
`

const ServiceResponse = `
$util.toJson($context.result)
`

func NewServiceResolver(mod module.Module, alias string, props *ServiceResolverProps) awsappsync.CfnResolver {

	escapedSchema := escapeSchema(props.Schema)

	request := fmt.Sprintf(ServiceRequest, *escapedSchema)
	response := ServiceResponse

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
