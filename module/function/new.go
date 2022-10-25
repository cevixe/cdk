package function

import (
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/lambda"
)

func NewFunction(mod module.Module, alias string, entry string) Function {

	fn := lambda.NewGolangFunction(
		mod,
		alias,
		&lambda.GolangFunctionProps{
			Entry: entry,
		},
	)

	return &functionImpl{
		module:   mod,
		name:     alias,
		resource: fn,
	}
}
