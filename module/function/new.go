package function

import (
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/lambda"
)

func NewFunction(mod module.Module, alias string, main string) Function {

	fn := lambda.NewGolangFunction(
		mod,
		alias,
		&lambda.GolangFunctionProps{
			Directory: mod.Location(),
			File:      main,
		},
	)

	return &functionImpl{
		module:   mod,
		name:     alias,
		resource: fn,
	}
}
