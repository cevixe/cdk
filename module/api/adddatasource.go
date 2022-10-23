package api

import (
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/cevixe/cdk/service/appsync"
)

type dataSourceImpl struct {
	name     string
	resource awsappsync.CfnDataSource
}

func (r *dataSourceImpl) Name() string {
	return r.name
}

func (r *dataSourceImpl) Resource() awsappsync.CfnDataSource {
	return r.resource
}

func (a *apiImpl) AddDataSource(alias string, props *DataSourceProps) DataSource {

	ds := &dataSourceImpl{name: alias}

	switch props.Type {
	case DSType_Mock:
		ds.resource = a.newMockDataSource(alias, props)
	case DSType_Function:
		ds.resource = a.newFunctionDataSource(alias, props)
	case DSType_StateStore:
		ds.resource = a.newStateStoreDataSource(alias, props)
	default:
		log.Fatal("unknown datasource type")
	}

	return ds
}

func (a *apiImpl) newMockDataSource(alias string, props *DataSourceProps) awsappsync.CfnDataSource {
	return appsync.NewDataSource(
		a.module,
		alias,
		&appsync.DataSourceProps{
			Api:  a.resource,
			Type: appsync.DSType_None,
		},
	)
}

func (a *apiImpl) newFunctionDataSource(alias string, props *DataSourceProps) awsappsync.CfnDataSource {
	return appsync.NewDataSource(
		a.module,
		alias,
		&appsync.DataSourceProps{
			Api:    a.resource,
			Role:   a.role,
			Type:   appsync.DSType_Lambda,
			Lambda: props.Function.Resource(),
		},
	)
}

func (a *apiImpl) newStateStoreDataSource(alias string, props *DataSourceProps) awsappsync.CfnDataSource {
	return appsync.NewDataSource(
		a.module,
		alias,
		&appsync.DataSourceProps{
			Api:   a.resource,
			Role:  a.role,
			Type:  appsync.DSType_Dynamodb,
			Table: props.StateStore.Resource(),
		},
	)
}
