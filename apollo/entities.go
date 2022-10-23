package apollo

import (
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/appsync"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
)

type EntitiesResolverProps struct {
	Api        awsappsync.CfnGraphQLApi `field:"required"`
	DataSource awsappsync.CfnDataSource `field:"required"`
	Table      awsdynamodb.Table        `field:"required"`
}

const EntitiesRequest = `
#set($ids = [])
#foreach($item in ${ctx.args.representations})
	#set($map = {})
	$util.qr($map.put("__typename", $util.dynamodb.toString($item.__typename)))
	$util.qr($map.put("id", $util.dynamodb.toString($item.id)))
	$util.qr($ids.add($map))
#end

{
	"version" : "2018-05-29",
	"operation" : "BatchGetItem",
	"tables" : {
    	"%s": {
    		"keys": $util.toJson($ids),
			"consistentRead": true
		}
    }
}
`

const EntitiesResponse = `
#if($ctx.error)
	$util.error($ctx.error.message, $ctx.error.type)
#end
$util.toJson($context.result)
`

func NewEntitiesResolver(mod module.Module, alias string, props *EntitiesResolverProps) awsappsync.CfnResolver {

	request := EntitiesRequest
	response := EntitiesResponse

	return appsync.NewResolver(
		mod,
		alias,
		&appsync.ResolverProps{
			Api:        props.Api,
			DataSource: props.DataSource,
			Type:       "Query",
			Field:      "_entities",
			Request:    jsii.String(request),
			Response:   jsii.String(response),
		},
	)
}
