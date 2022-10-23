package iam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

func NewLambdaInvokePol(lambda awslambda.Function) awsiam.PolicyStatement {

	return awsiam.NewPolicyStatement(
		&awsiam.PolicyStatementProps{
			Effect: awsiam.Effect_ALLOW,
			Actions: &[]*string{
				jsii.String("lambda:invokeFunction"),
			},
			Resources: &[]*string{
				lambda.FunctionArn(),
			},
		},
	)
}
