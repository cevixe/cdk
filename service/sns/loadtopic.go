package sns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

func LoadTopic(mod module.Module, alias string, topicArn string) awssns.ITopic {

	name := naming.NewName(mod, naming.ResType_SNSTopic, alias)

	return awssns.Topic_FromTopicArn(
		mod.Resource(),
		name.Logical(),
		jsii.String(topicArn),
	)
}
