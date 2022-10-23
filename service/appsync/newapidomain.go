package appsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

type ApiDomainProps struct {
	Api         awsappsync.CfnGraphQLApi          `field:"required"`
	Domain      string                            `field:"required"`
	Certificate awscertificatemanager.Certificate `field:"required"`
}

func NewApiDomain(mod module.Module, alias string, props *ApiDomainProps) awsappsync.CfnDomainName {

	domainName := naming.NewName(mod, naming.ResType_GraphQLDomain, alias)

	domain := awsappsync.NewCfnDomainName(
		mod.Resource(),
		domainName.Logical(),
		&awsappsync.CfnDomainNameProps{
			DomainName:     jsii.String(props.Domain),
			CertificateArn: props.Certificate.CertificateArn(),
		},
	)

	linkName := naming.NewName(mod, naming.ResType_GraphQLDomainLink, alias)

	link := awsappsync.NewCfnDomainNameApiAssociation(
		mod.Resource(),
		linkName.Logical(),
		&awsappsync.CfnDomainNameApiAssociationProps{
			ApiId:      props.Api.AttrApiId(),
			DomainName: jsii.String(props.Domain),
		},
	)

	link.AddDependsOn(domain)

	return domain
}
