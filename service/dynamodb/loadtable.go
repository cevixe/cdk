package dynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

func LoadTable(mod module.Module, alias string, tableArn string) awsdynamodb.ITable {

	name := naming.NewName(mod, naming.ResType_DynamodbTable, alias)

	return awsdynamodb.Table_FromTableAttributes(
		mod.Resource(),
		name.Logical(),
		&awsdynamodb.TableAttributes{
			TableArn: jsii.String(tableArn),
		},
	)
}
