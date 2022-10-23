package domain

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/module/api"
	"github.com/cevixe/cdk/module/handler"
	"github.com/cevixe/cdk/module/objectstore"
	"github.com/cevixe/cdk/module/statestore"
)

type Domain interface {
	module.Module

	Api() api.Api

	StateStoreDS() api.DataSource

	StateStore() statestore.StateStore
	ObjectStore() objectstore.ObjectStore

	NewHandler(alias string, props *handler.HandlerProps) handler.Handler
	NewResolver(alias string, props *api.ResolverProps) api.Resolver
}

type Context struct {
	zone awsroute53.IHostedZone

	advancedBus awssns.ITopic
	standardBus awssns.ITopic

	advancedCdc   awslambda.IFunction
	standardCdc   awslambda.IFunction
	linkGenerator awslambda.IFunction

	eventStore   awsdynamodb.ITable
	commandStore awsdynamodb.ITable
}
