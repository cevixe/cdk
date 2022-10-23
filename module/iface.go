package module

import "github.com/aws/constructs-go/constructs/v10"

type Module interface {
	App() string
	Name() string
	Location() string
	DependsOn(mod ...Module)
	Export(alias string, value string)
	Resource() constructs.Construct
}
