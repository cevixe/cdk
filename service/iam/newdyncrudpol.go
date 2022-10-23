package iam

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
)

func NewDynCrudPol(table awsdynamodb.Table) awsiam.PolicyStatement {

	return awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: &[]*string{
			jsii.String("dynamodb:GetItem"),
			jsii.String("dynamodb:DeleteItem"),
			jsii.String("dynamodb:PutItem"),
			jsii.String("dynamodb:Scan"),
			jsii.String("dynamodb:Query"),
			jsii.String("dynamodb:UpdateItem"),
			jsii.String("dynamodb:BatchWriteItem"),
			jsii.String("dynamodb:BatchGetItem"),
			jsii.String("dynamodb:DescribeTable"),
			jsii.String("dynamodb:ConditionCheckItem"),
		},
		Resources: &[]*string{
			table.TableArn(),
			jsii.String(fmt.Sprintf("%s/index/*", *table.TableArn())),
		},
	})
}
