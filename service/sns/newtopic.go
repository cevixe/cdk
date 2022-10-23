package sns

import (
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

type TopicType uint8

const (
	TopicType_Standard TopicType = 0
	TopicType_FIFO     TopicType = 1
)

type TopicProps struct {
	Type TopicType `field:"required"`
}

func NewTopic(mod module.Module, alias string, props *TopicProps) awssns.Topic {

	if props.Type == TopicType_FIFO && !strings.HasSuffix(alias, ".fifo") {
		alias = alias + ".fifo"
	}

	name := naming.NewName(mod, naming.ResType_SQSQueue, alias)

	awsprops := &awssns.TopicProps{
		TopicName: name.Physical(),
	}

	if props.Type == TopicType_FIFO {
		awsprops.Fifo = jsii.Bool(true)
		awsprops.ContentBasedDeduplication = jsii.Bool(true)
	}

	return awssns.NewTopic(mod.Resource(), name.Logical(), awsprops)
}
