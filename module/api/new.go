package api

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/cevixe/cdk/common/file"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/acm"
	"github.com/cevixe/cdk/service/appsync"
	"github.com/cevixe/cdk/service/route53"
)

type ApiProps struct {
	Domain string                 `field:"required"`
	Zone   awsroute53.IHostedZone `field:"required"`
}

func NewApi(mod module.Module, alias string, props *ApiProps) Api {

	domainName := fmt.Sprintf("%s.%s", mod.Name(), props.Domain)
	api := appsync.NewApi(mod, alias)

	schemaContent := file.GetFileContent(
		fmt.Sprintf("%s/cdk/schema.gql", mod.Location()))

	schema := appsync.NewSchema(mod, alias, &appsync.SchemaProps{
		Api:        api,
		Definition: &schemaContent,
	})

	role := appsync.NewApiRole(mod, alias)

	certificate := acm.NewCertificate(
		mod,
		alias,
		&acm.CertificateProps{
			Zone:   props.Zone,
			Domain: domainName,
		},
	)

	apiDomain := appsync.NewApiDomain(
		mod,
		alias,
		&appsync.ApiDomainProps{
			Api:         api,
			Domain:      domainName,
			Certificate: certificate,
		},
	)

	route53.NewCnameRecord(
		mod,
		alias,
		&route53.CnameRecordProps{
			Zone:   props.Zone,
			Record: mod.Name(),
			Domain: *apiDomain.AttrAppSyncDomainName(),
		},
	)

	return &apiImpl{
		module:   mod,
		name:     alias,
		record:   mod.Name(),
		domain:   props.Domain,
		schema:   schema,
		role:     role,
		resource: api,
	}
}
