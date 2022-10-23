package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

func LoadLambda(mod module.Module, alias string, lambdaArn string, lambdaRole awsiam.IRole) awslambda.IFunction {

	name := naming.NewName(mod, naming.ResType_Lambda, alias)

	return awslambda.Function_FromFunctionAttributes(
		mod.Resource(),
		name.Logical(),
		&awslambda.FunctionAttributes{
			FunctionArn: jsii.String(lambdaArn),
			Role:        lambdaRole,
		},
	)
}
