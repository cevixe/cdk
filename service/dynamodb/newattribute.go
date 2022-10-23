package dynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
)

type Attribute = awsdynamodb.Attribute

func NewAttribute(name string, typ awsdynamodb.AttributeType) *Attribute {
	return &awsdynamodb.Attribute{
		Name: jsii.String(name),
		Type: typ,
	}
}
