package application

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	loc "github.com/cevixe/app/pkg/location"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/module/bus"
	"github.com/cevixe/cdk/module/eventstore"
	"github.com/cevixe/cdk/module/function"
	"github.com/cevixe/cdk/module/handler"
	"github.com/cevixe/cdk/module/objectstore"
	"github.com/cevixe/cdk/module/statestore"
	"github.com/cevixe/cdk/naming"
	"github.com/cevixe/cdk/service/iam"
	"github.com/cevixe/cdk/service/sns"
	"github.com/cevixe/cdk/service/sqs"
)

type ApplicationProps struct {
	Location string
	Domains  []string
}

func NewApplication(scope constructs.Construct, app string, props *ApplicationProps) {
	newCore(scope, app, props.Location)

	for _, dom := range props.Domains {
		mod := newStoreModule(scope, app, dom)
		ss := statestore.NewStateStore(mod, dom)
		os := objectstore.NewObjectStore(mod, dom)

		mod.Export(StateStoreArn, *ss.Resource().TableArn())
		mod.Export(StateStoreName, *ss.Resource().TableName())

		mod.Export(ObjectStoreArn, *os.Resource().BucketArn())
		mod.Export(ObjectStoreName, *os.Resource().BucketName())
	}
}

func newCore(scope constructs.Construct, app string, location string) module.Module {
	alias := "core"
	mod := newModule(scope, app, alias, location)

	advancedbus := bus.NewBus(mod, "advancedbus", &bus.BusProps{Type: bus.BusType_Advanced})
	standardbus := bus.NewBus(mod, "standardbus", &bus.BusProps{Type: bus.BusType_Standard})

	commandstore := statestore.NewStateStore(mod, "commandstore")
	eventstore := eventstore.NewEventStore(mod, "eventstore")

	advancedcdc := function.NewFunction(mod, "advancedcdc", loc.AdvancedCdc)
	standardcdc := function.NewFunction(mod, "standardcdc", loc.StandardCdc)

	eventhandler := handler.NewHandler(mod, "eventhandler",
		&handler.HandlerProps{
			Type:   handler.HandlerType_Advanced,
			Events: &[]string{},
			Main:   loc.EventHandler,
		},
	)

	advancedcdc.Resource().AddToRolePolicy(iam.NewSNSPublishPol(advancedbus.Resource()))
	standardcdc.Resource().AddToRolePolicy(iam.NewSNSPublishPol(standardbus.Resource()))
	stream := awslambdaeventsources.NewDynamoEventSource(commandstore.Resource(),
		&awslambdaeventsources.DynamoEventSourceProps{
			Enabled:          jsii.Bool(true),
			BatchSize:        jsii.Number(10),
			StartingPosition: awslambda.StartingPosition_TRIM_HORIZON,
		})
	advancedcdc.Resource().AddEventSource(stream)
	standardcdc.Resource().AddEventSource(stream)

	eventhandler.Resource().AddToRolePolicy(iam.NewDynCrudPol(eventstore.Resource()))
	sns.NewSubscriptions(mod, eventhandler.Name(), &sns.SubProps{
		Topic:    advancedbus.Resource(),
		Function: eventhandler.Resource(),
		Filters:  &[]*bus.Filter{bus.NewFilter("event")},
		Queue:    sqs.NewQueue(mod, eventhandler.Name(), &sqs.QueueProps{Type: sqs.QueueType_FIFO}),
	})

	mod.Export(AdvancedBusArn, *advancedbus.Resource().TopicArn())
	mod.Export(AdvancedBusName, *advancedbus.Resource().TopicName())

	mod.Export(StandardBusArn, *standardbus.Resource().TopicArn())
	mod.Export(StandardBusName, *standardbus.Resource().TopicName())

	mod.Export(CommandStoreArn, *commandstore.Resource().TableArn())
	mod.Export(CommandStoreName, *commandstore.Resource().TableName())

	mod.Export(EventStoreArn, *eventstore.Resource().TableArn())
	mod.Export(EventStoreName, *eventstore.Resource().TableName())

	mod.Export(AdvancedCdcArn, *advancedcdc.Resource().FunctionArn())
	mod.Export(AdvancedCdcName, *advancedcdc.Resource().FunctionName())
	mod.Export(AdvancedCdcRole, *advancedcdc.Resource().Role().RoleArn())

	mod.Export(StandardCdcArn, *standardcdc.Resource().FunctionArn())
	mod.Export(StandardCdcName, *standardcdc.Resource().FunctionName())
	mod.Export(StandardCdcRole, *standardcdc.Resource().Role().RoleArn())

	mod.Export(EventHandlerArn, *eventhandler.Resource().FunctionArn())
	mod.Export(EventHandlerName, *eventhandler.Resource().FunctionName())
	mod.Export(EventHandlerRole, *eventhandler.Resource().Role().RoleArn())

	return mod
}

func newStoreModule(scope constructs.Construct, app string, alias string) module.Module {

	mod := &moduleImpl{
		app:      app,
		name:     alias,
		location: ".",
	}

	name := naming.NewName(mod, naming.ResType_Stack, "store")

	mod.resource = awscdk.NewStack(
		scope,
		&alias,
		&awscdk.StackProps{
			StackName: name.Physical(),
		},
	)

	return mod
}

func newModule(scope constructs.Construct, app string, alias string, location string) module.Module {

	mod := &moduleImpl{
		app:      app,
		name:     alias,
		location: location,
	}

	name := naming.NewName(mod, naming.ResType_Stack, "module")

	mod.resource = awscdk.NewStack(
		scope,
		&alias,
		&awscdk.StackProps{
			StackName: name.Physical(),
		},
	)

	return mod
}
