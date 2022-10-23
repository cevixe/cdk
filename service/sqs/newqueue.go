package sqs

import (
	"strings"

	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/jsii-runtime-go"
)

type QueueType uint8

const (
	QueueType_Standard QueueType = 0
	QueueType_FIFO     QueueType = 1
)

type QueueProps struct {
	Type QueueType `field:"required"`
}

func NewQueue(mod module.Module, alias string, props *QueueProps) awssqs.Queue {

	if props.Type == QueueType_FIFO && !strings.HasSuffix(alias, ".fifo") {
		alias = alias + ".fifo"
	}

	name := naming.NewName(mod, naming.ResType_SQSQueue, alias)

	awsprops := &awssqs.QueueProps{
		QueueName:       name.Physical(),
		RemovalPolicy:   awscdk.RemovalPolicy_DESTROY,
		RetentionPeriod: awscdk.Duration_Days(jsii.Number(14)),
	}

	if props.Type == QueueType_FIFO {
		awsprops.Fifo = jsii.Bool(true)
		awsprops.ContentBasedDeduplication = jsii.Bool(true)
		awsprops.DeduplicationScope = awssqs.DeduplicationScope_MESSAGE_GROUP
		awsprops.FifoThroughputLimit = awssqs.FifoThroughputLimit_PER_MESSAGE_GROUP_ID
	}

	return awssqs.NewQueue(mod.Resource(), name.Logical(), awsprops)
}
