package project

import (
	"log"
	"os"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cevixe/cdk/application"
	spec20221023 "github.com/cevixe/cdk/spec/20221023"
	"gopkg.in/yaml.v3"
)

func loadProject20221023(scope constructs.Construct, fileBytes *[]byte) {

	file := &spec20221023.File{}
	err := yaml.Unmarshal(*fileBytes, file)
	if err != nil {
		log.Fatalf("invalid configuration file structure: %v", err)
	}

	location := os.Getenv("CEVIXE_APP_DIR")
	if location == "" {
		log.Fatalf("CEVIXE_APP_DIR not configured")
	}

	application.NewApplication(scope, file.Project.Name, &application.ApplicationProps{
		Location: location,
		Domains:  *file.Project.Domains,
	})
}
