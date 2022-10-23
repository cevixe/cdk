package function

import (
	"fmt"

	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/lambda"
)

func NewFunction(mod module.Module, alias string) Function {

	fn := lambda.NewGolangFunction(
		mod,
		alias,
		&lambda.GolangFunctionProps{
			Directory: fmt.Sprintf("%s/app", mod.Location()),
			File:      fmt.Sprintf("cmd/%s/main.go", alias),
		},
	)

	return &functionImpl{
		module:   mod,
		name:     alias,
		resource: fn,
	}
}
