package lambda

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/jsii-runtime-go"
)

type GolangFunctionProps struct {
	Directory string
	File      string
}

func NewGolangFunction(mod module.Module, alias string, props *GolangFunctionProps) awslambda.Function {

	name := naming.NewName(mod, naming.ResType_Lambda, alias)
	role := NewFunctionRole(mod, alias)

	return awslambda.NewFunction(
		mod.Resource(),
		name.Logical(),
		&awslambda.FunctionProps{
			FunctionName: name.Physical(),
			Architecture: awslambda.Architecture_X86_64(),
			Tracing:      awslambda.Tracing_ACTIVE,
			Runtime:      awslambda.Runtime_GO_1_X(),
			MemorySize:   jsii.Number(256),
			Code:         newGolangFunctionCode(props.Directory, props.File),
			Handler:      jsii.String("handler"),
			Role:         role,
		},
	)
}

func newGolangFunctionCode(directory string, file string) awslambda.Code {

	env := "CGO_ENABLED=0 GOOS=linux GOARCH=amd64"
	config := "-ldflags=\"-w -s\""
	artifact := "/asset-output/handler"
	command := fmt.Sprintf("pwd && ls -laR . && %s go build %s -o %s %s", env, config, artifact, file)

	fmt.Println("====================================================")
	fmt.Printf("Directory: %s\n", directory)
	fmt.Printf("File: %s\n", file)
	fmt.Printf("Command: %s\n", command)

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
	fmt.Println("====================================================")
	return awslambda.Code_FromAsset(
		jsii.String(directory),
		&awss3assets.AssetOptions{
			Bundling: &awscdk.BundlingOptions{
				User:  jsii.String("root"),
				Image: awslambda.Runtime_GO_1_X().BundlingImage(),
				Command: &[]*string{
					jsii.String("bash"),
					jsii.String("-c"),
					jsii.String(command),
				},
			},
		},
	)
}
