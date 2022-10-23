package api

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/cevixe/cdk/common/file"
	"github.com/cevixe/cdk/service/appsync"
)

type resolverImpl struct {
	name     string
	resource awsappsync.CfnResolver
}

func (r *resolverImpl) Name() string {
	return r.name
}

func (r *resolverImpl) Resource() awsappsync.CfnResolver {
	return r.resource
}

func (a *apiImpl) AddResolver(alias string, props *ResolverProps) Resolver {

	location := a.module.Location()

	requestLocation := fmt.Sprintf("%s/cdk/templates/%s/request.vtl", location, alias)
	request := file.GetFileContent(requestLocation)

	responseLocation := fmt.Sprintf("%s/cdk/templates/%s/response.vtl", location, alias)
	response := file.GetFileContent(responseLocation)

	resource := appsync.NewResolver(
		a.module,
		alias,
		&appsync.ResolverProps{
			Api:        a.Resource(),
			DataSource: props.DataSource.Resource(),
			Type:       props.Type,
			Field:      props.Field,
			Request:    &request,
			Response:   &response,
		},
	)

	resource.AddDependsOn(a.schema)
	return &resolverImpl{
		name:     alias,
		resource: resource,
	}
}
