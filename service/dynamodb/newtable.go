package dynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

type TableProps struct {
	Key           *Key                   `field:"required"`
	LocalIndexes  *map[string]*Attribute `field:"optional"`
	GlobalIndexes *map[string]*Key       `field:"optional"`
}

func NewTable(mod module.Module, alias string, props *TableProps) awsdynamodb.Table {

	name := naming.NewName(mod, naming.ResType_DynamodbTable, alias)

	table := awsdynamodb.NewTable(
		mod.Resource(),
		name.Logical(),
		&awsdynamodb.TableProps{
			TableName:           name.Physical(),
			PartitionKey:        props.Key.PartitionKey,
			SortKey:             props.Key.SortKey,
			Encryption:          awsdynamodb.TableEncryption_AWS_MANAGED,
			BillingMode:         awsdynamodb.BillingMode_PAY_PER_REQUEST,
			Stream:              awsdynamodb.StreamViewType_NEW_AND_OLD_IMAGES,
			TimeToLiveAttribute: jsii.String("__ttl"),
			RemovalPolicy:       awscdk.RemovalPolicy_DESTROY,
		},
	)

	if props.LocalIndexes != nil {
		for name, sortKey := range *props.LocalIndexes {
			table.AddLocalSecondaryIndex(&awsdynamodb.LocalSecondaryIndexProps{
				IndexName:      jsii.String(name),
				ProjectionType: awsdynamodb.ProjectionType_ALL,
				SortKey:        sortKey,
			})
		}
	}

	if props.GlobalIndexes != nil {
		for name, composedKey := range *props.GlobalIndexes {
			table.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
				IndexName:      jsii.String(name),
				ProjectionType: awsdynamodb.ProjectionType_ALL,
				PartitionKey:   composedKey.PartitionKey,
				SortKey:        composedKey.SortKey,
			})
		}
	}

	return table
}
