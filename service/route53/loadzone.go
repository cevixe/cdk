package route53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

func LoadZone(mod module.Module, alias string, zoneId string, zoneName string) awsroute53.IHostedZone {

	name := naming.NewName(mod, naming.ResType_Route53Zone, alias)

	return awsroute53.PublicHostedZone_FromHostedZoneAttributes(
		mod.Resource(),
		name.Logical(),
		&awsroute53.HostedZoneAttributes{
			ZoneName:     jsii.String(zoneName),
			HostedZoneId: jsii.String(zoneId),
		},
	)
}
