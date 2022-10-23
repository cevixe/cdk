package naming

import (
	"fmt"
	"strings"

	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Name interface {
	Physical() *string
	Logical() *string
}

type nameImpl struct {
	physical *string
	logical  *string
}

func (n *nameImpl) Physical() *string {
	return n.physical
}

func (n *nameImpl) Logical() *string {
	return n.logical
}

func NewName(mod module.Module, resType ResType, alias string) Name {

	physicalFormat := "%s-%s-%s-%s"
	physicalName := fmt.Sprintf(
		physicalFormat,
		resType,
		strings.ToLower(mod.App()),
		strings.ToLower(mod.Name()),
		strings.ToLower(alias),
	)

	caser := cases.Title(language.English)
	logicalFormat := "%s%s"
	logicalName := fmt.Sprintf(
		logicalFormat,
		caser.String(string(resType)),
		caser.String(alias),
	)

	return &nameImpl{
		physical: jsii.String(physicalName),
		logical:  jsii.String(logicalName),
	}
}
