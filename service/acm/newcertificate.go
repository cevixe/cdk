package acm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

type CertificateProps struct {
	Zone   awsroute53.IHostedZone `field:"required"`
	Domain string                 `field:"required"`
}

func NewCertificate(mod module.Module, alias string, props *CertificateProps) awscertificatemanager.Certificate {

	name := naming.NewName(mod, naming.ResType_ACMCertificate, alias)

	return awscertificatemanager.NewCertificate(
		mod.Resource(),
		name.Logical(),
		&awscertificatemanager.CertificateProps{
			DomainName: jsii.String(props.Domain),
			Validation: awscertificatemanager.CertificateValidation_FromDns(props.Zone),
		},
	)
}
