package route53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

type CnameRecordProps struct {
	Zone   awsroute53.IHostedZone `field:"required"`
	Record string                 `field:"required"`
	Domain string                 `field:"required"`
}

func NewCnameRecord(mod module.Module, alias string, props *CnameRecordProps) awsroute53.CnameRecord {

	name := naming.NewName(mod, naming.ResType_Route53Record, alias)

	return awsroute53.NewCnameRecord(
		mod.Resource(),
		name.Logical(),
		&awsroute53.CnameRecordProps{
			Zone:       props.Zone,
			RecordName: jsii.String(props.Record),
			DomainName: jsii.String(props.Domain),
		},
	)
}
