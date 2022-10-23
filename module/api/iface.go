package api

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/module/function"
	"github.com/cevixe/cdk/module/statestore"
)

type Api interface {
	Module() module.Module
	Name() string
	Domain() string
	Schema() string
	AddResolver(alias string, props *ResolverProps) Resolver
	AddDataSource(alias string, props *DataSourceProps) DataSource
	Resource() awsappsync.CfnGraphQLApi
}

type ResolverProps struct {
	DataSource DataSource `field:"required"`
	Type       string     `field:"required"`
	Field      string     `field:"required"`
}

type Resolver interface {
	Name() string
	Resource() awsappsync.CfnResolver
}

type DSType uint8

const (
	DSType_Mock       DSType = 0
	DSType_Function   DSType = 1
	DSType_StateStore DSType = 2
)

type DataSourceProps struct {
	Type       DSType                `field:"required"`
	Function   function.Function     `field:"optional"`
	StateStore statestore.StateStore `field:"optional"`
}

type DataSource interface {
	Name() string
	Resource() awsappsync.CfnDataSource
}
