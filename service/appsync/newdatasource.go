package appsync

import (
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
	"github.com/cevixe/cdk/service/iam"
)

type DSType uint8

const (
	DSType_None     DSType = 0
	DSType_Lambda   DSType = 1
	DSType_Dynamodb DSType = 2
)

type DataSourceProps struct {
	Api    awsappsync.CfnGraphQLApi `field:"required"`
	Type   DSType                   `field:"required"`
	Role   awsiam.Role              `field:"optional"`
	Table  awsdynamodb.Table        `field:"optional"`
	Lambda awslambda.Function       `field:"optional"`
}

func NewDataSource(mod module.Module, alias string, props *DataSourceProps) awsappsync.CfnDataSource {

	switch props.Type {
	case DSType_None:
		return newNoneDS(mod, alias, props)
	case DSType_Lambda:
		return newLambdaDS(mod, alias, props)
	case DSType_Dynamodb:
		return newDynamodbDS(mod, alias, props)
	default:
		log.Fatal("unknown appsync datasource type")
	}
	return nil
}

func newNoneDS(mod module.Module, alias string, props *DataSourceProps) awsappsync.CfnDataSource {

	name := naming.NewName(mod, naming.ResType_GraphQLDataSource, alias)

	return awsappsync.NewCfnDataSource(
		mod.Resource(),
		name.Logical(),
		&awsappsync.CfnDataSourceProps{
			Name:  name.Logical(),
			ApiId: props.Api.AttrApiId(),
			Type:  jsii.String("NONE"),
		},
	)
}

func newLambdaDS(mod module.Module, alias string, props *DataSourceProps) awsappsync.CfnDataSource {

	name := naming.NewName(mod, naming.ResType_GraphQLDataSource, alias)

	props.Role.AddToPolicy(iam.NewLambdaInvokePol(props.Lambda))

	return awsappsync.NewCfnDataSource(
		mod.Resource(),
		name.Logical(),
		&awsappsync.CfnDataSourceProps{
			Name:  name.Logical(),
			ApiId: props.Api.AttrApiId(),
			Type:  jsii.String("AWS_LAMBDA"),
			LambdaConfig: &awsappsync.CfnDataSource_LambdaConfigProperty{
				LambdaFunctionArn: props.Lambda.FunctionArn(),
			},
			ServiceRoleArn: props.Role.RoleArn(),
		},
	)
}

func newDynamodbDS(mod module.Module, alias string, props *DataSourceProps) awsappsync.CfnDataSource {

	name := naming.NewName(mod, naming.ResType_GraphQLDataSource, alias)

	props.Role.AddToPolicy(iam.NewDynCrudPol(props.Table))

	return awsappsync.NewCfnDataSource(
		mod.Resource(),
		name.Logical(),
		&awsappsync.CfnDataSourceProps{
			Name:           name.Logical(),
			ApiId:          props.Api.AttrApiId(),
			Type:           jsii.String("AMAZON_DYNAMODB"),
			ServiceRoleArn: props.Role.RoleArn(),
			DynamoDbConfig: awsappsync.CfnDataSource_DynamoDBConfigProperty{
				AwsRegion: props.Table.Stack().Region(),
				TableName: props.Table.TableName(),
			},
		},
	)
}
