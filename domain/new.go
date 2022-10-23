package domain

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cevixe/cdk/apollo"
	"github.com/cevixe/cdk/core"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/module/api"
	"github.com/cevixe/cdk/module/objectstore"
	"github.com/cevixe/cdk/module/statestore"
	"github.com/cevixe/cdk/naming"
	"github.com/cevixe/cdk/service/dynamodb"
	"github.com/cevixe/cdk/service/iam"
	"github.com/cevixe/cdk/service/lambda"
	"github.com/cevixe/cdk/service/route53"
	"github.com/cevixe/cdk/service/sns"
)

func NewDomain(scope constructs.Construct, app string, alias string) Domain {

	mod := newModule(scope, app, alias)

	mod.statestore = statestore.NewStateStore(mod, "statestore")
	mod.objectstore = objectstore.NewObjectStore(mod, "objectstore")

	mod.api = api.NewApi(mod, "api", &api.ApiProps{
		Domain: fmt.Sprintf("%s.%s", strings.ToLower(alias), mod.domain),
		Zone:   mod.context.zone,
	})

	mod.mockds = mod.api.AddDataSource("mock", &api.DataSourceProps{
		Type: api.DSType_Mock,
	})

	mod.statestoreds = mod.api.AddDataSource("statestore", &api.DataSourceProps{
		Type:       api.DSType_StateStore,
		StateStore: mod.statestore,
	})

	apollo.NewEntitiesResolver(mod, "entities", &apollo.EntitiesResolverProps{
		Api:        mod.api.Resource(),
		DataSource: mod.statestoreds.Resource(),
		Table:      mod.statestore.Resource(),
	})

	schema := mod.api.Schema()
	apollo.NewServiceResolver(mod, "service", &apollo.ServiceResolverProps{
		Api:        mod.api.Resource(),
		DataSource: mod.mockds.Resource(),
		Schema:     &schema,
	})

	return mod
}

func newModule(scope constructs.Construct, app string, alias string) *domainImpl {

	location := fmt.Sprintf("modules/%s", strings.ToLower(alias))

	mod := &domainImpl{
		app:      app,
		name:     alias,
		location: location,
		domain:   os.Getenv("AWS_DOMAIN"),
	}

	name := naming.NewName(mod, naming.ResType_Stack, "module")

	mod.resource = awscdk.NewStack(
		scope,
		&alias,
		&awscdk.StackProps{
			StackName: name.Physical(),
		},
	)

	ctx := &Context{}

	ctx.zone = route53.LoadZone(mod, "dns", os.Getenv("AWS_ZONE"), os.Getenv("AWS_DOMAIN"))

	commandstorearn := importVar(mod, core.CommandStoreArn)
	//commandstorename := importVar(mod, core.CommandStoreName)
	ctx.commandStore = dynamodb.LoadTable(mod, "commandstore", *commandstorearn)

	eventstorearn := importVar(mod, core.EventStoreArn)
	//eventstorename := importVar(mod, core.EventStoreName)
	ctx.eventStore = dynamodb.LoadTable(mod, "eventstore", *eventstorearn)

	advancedbusarn := importVar(mod, core.AdvancedBusArn)
	ctx.advancedBus = sns.LoadTopic(mod, "advancedbus", *advancedbusarn)

	standardbusarn := importVar(mod, core.StandardBusArn)
	ctx.standardBus = sns.LoadTopic(mod, "standardbus", *standardbusarn)

	standardcdcarn := importVar(mod, core.StandardCdcArn)
	standardcdcrole := iam.LoadServiceRole(mod, "standardcdc", *importVar(mod, core.StandardCdcRole))
	ctx.standardCdc = lambda.LoadLambda(mod, "standardcdc", *standardcdcarn, standardcdcrole)

	advancedcdcarn := importVar(mod, core.AdvancedCdcArn)
	advancedcdcrole := iam.LoadServiceRole(mod, "advancedcdc", *importVar(mod, core.AdvancedCdcRole))
	ctx.advancedCdc = lambda.LoadLambda(mod, "advancedcdc", *advancedcdcarn, advancedcdcrole)

	mod.context = ctx

	return mod
}

func importVar(mod module.Module, alias string) *string {
	name := fmt.Sprintf("%s-%s-%s-%s", naming.ResType_Output, mod.App(), "core", alias)
	return awscdk.Fn_ImportValue(&name)
}
