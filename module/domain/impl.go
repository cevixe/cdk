package domain

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/module/api"
	"github.com/cevixe/cdk/module/handler"
	"github.com/cevixe/cdk/module/objectstore"
	"github.com/cevixe/cdk/module/statestore"
	"github.com/cevixe/cdk/naming"
)

type domainImpl struct {
	app          string
	name         string
	location     string
	domain       string
	api          api.Api
	mockds       api.DataSource
	statestoreds api.DataSource
	statestore   statestore.StateStore
	objectstore  objectstore.ObjectStore
	resource     constructs.Construct
	context      *Context
}

func (m *domainImpl) App() string {
	return m.app
}

func (m *domainImpl) Name() string {
	return m.name
}

func (m *domainImpl) Location() string {
	return m.location
}

func (m *domainImpl) Resource() constructs.Construct {
	return m.resource
}

func (m *domainImpl) Api() api.Api {
	return m.api
}

func (m *domainImpl) StateStoreDS() api.DataSource {
	return m.statestoreds
}

func (m *domainImpl) StateStore() statestore.StateStore {
	return m.statestore
}

func (m *domainImpl) ObjectStore() objectstore.ObjectStore {
	return m.objectstore
}

func (m *domainImpl) Export(alias string, value string) {

	name := naming.NewName(m, naming.ResType_Output, alias)

	awscdk.NewCfnOutput(
		m.Resource(),
		name.Logical(),
		&awscdk.CfnOutputProps{
			Value:      jsii.String(value),
			ExportName: name.Physical(),
		},
	)
}

func (m *domainImpl) NewHandler(alias string, props *handler.HandlerProps) handler.Handler {
	return handler.NewHandler(m, alias, props)
}

func (m *domainImpl) NewResolver(alias string, props *api.ResolverProps) api.Resolver {
	return m.Api().AddResolver(alias, props)
}

func (m *domainImpl) DependsOn(mod ...module.Module) {
	for _, item := range mod {
		m.resource.Node().AddDependency(item.Resource())
	}
}
