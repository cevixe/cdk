package eventstore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/dynamodb"
)

func NewEventStore(mod module.Module, alias string) EventStore {

	table := dynamodb.NewTable(mod, alias, &dynamodb.TableProps{
		Key: &dynamodb.Key{
			PartitionKey: dynamodb.NewAttribute("source", awsdynamodb.AttributeType_STRING),
			SortKey:      dynamodb.NewAttribute("id", awsdynamodb.AttributeType_STRING),
		},
		GlobalIndexes: &map[string]*dynamodb.Key{
			"by-type": {
				PartitionKey: dynamodb.NewAttribute("type", awsdynamodb.AttributeType_STRING),
				SortKey:      dynamodb.NewAttribute("time", awsdynamodb.AttributeType_STRING),
			},
			"by-user": {
				PartitionKey: dynamodb.NewAttribute("user", awsdynamodb.AttributeType_STRING),
				SortKey:      dynamodb.NewAttribute("time", awsdynamodb.AttributeType_STRING),
			},
			"by-transaction": {
				PartitionKey: dynamodb.NewAttribute("transaction", awsdynamodb.AttributeType_STRING),
				SortKey:      dynamodb.NewAttribute("time", awsdynamodb.AttributeType_STRING),
			},
		},
	})

	return &eventStoreImpl{
		module:   mod,
		name:     alias,
		resource: table,
	}
}
