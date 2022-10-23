package iam

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/jsii-runtime-go"
)

func NewS3CrudPol(bucket awss3.Bucket) awsiam.PolicyStatement {

	return awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: &[]*string{
			jsii.String("s3:GetObject"),
			jsii.String("s3:ListBucket"),
			jsii.String("s3:GetBucketLocation"),
			jsii.String("s3:GetObjectVersion"),
			jsii.String("s3:PutObject"),
			jsii.String("s3:PutObjectAcl"),
			jsii.String("s3:GetLifecycleConfiguration"),
			jsii.String("s3:PutLifecycleConfiguration"),
			jsii.String("s3:DeleteObject"),
		},
		Resources: &[]*string{
			bucket.BucketArn(),
			jsii.String(fmt.Sprintf("%s/*", *bucket.BucketArn())),
		},
	})
}
