package api

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/cevixe/cdk/module"
)

type apiImpl struct {
	module   module.Module
	name     string
	record   string
	domain   string
	schema   awsappsync.CfnGraphQLSchema
	role     awsiam.Role
	resource awsappsync.CfnGraphQLApi
}

func (a *apiImpl) Module() module.Module {
	return a.module
}

func (a *apiImpl) Name() string {
	return a.name
}

func (a *apiImpl) Schema() string {
	return *a.schema.Definition()
}

func (a *apiImpl) Domain() string {
	return fmt.Sprintf("%s.%s", a.record, a.domain)
}

func (a *apiImpl) Resource() awsappsync.CfnGraphQLApi {
	return a.resource
}
