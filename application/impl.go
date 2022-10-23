package application

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

type moduleImpl struct {
	app      string
	name     string
	location string
	resource constructs.Construct
}

func (m *moduleImpl) App() string {
	return m.app
}

func (m *moduleImpl) Name() string {
	return m.name
}

func (m *moduleImpl) Location() string {
	return m.location
}

func (m *moduleImpl) Resource() constructs.Construct {
	return m.resource
}

func (m *moduleImpl) Export(alias string, value string) {

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

func (m *moduleImpl) DependsOn(mod ...module.Module) {
	for _, item := range mod {
		m.resource.Node().AddDependency(item.Resource())
	}
}
