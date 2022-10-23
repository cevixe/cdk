package statestore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/dynamodb"
)

func NewStateStore(mod module.Module, alias string) StateStore {

	table := dynamodb.NewTable(mod, alias, &dynamodb.TableProps{
		Key: &dynamodb.Key{
			PartitionKey: dynamodb.NewAttribute("__typename", awsdynamodb.AttributeType_STRING),
			SortKey:      dynamodb.NewAttribute("id", awsdynamodb.AttributeType_STRING),
		},
		LocalIndexes: &map[string]*dynamodb.Attribute{
			"by-time": dynamodb.NewAttribute("__time", awsdynamodb.AttributeType_STRING),
		},
	})

	return &stateStoreImpl{
		module:   mod,
		name:     alias,
		resource: table,
	}
}
