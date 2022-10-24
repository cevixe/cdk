package project

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cevixe/cdk/common/file"
	"gopkg.in/yaml.v3"
)

func Load(scope constructs.Construct) {

	dir := os.Getenv("CEVIXE_WORKSPACE")
	spec := readSpecFile(dir)
	genericMap := make(map[string]interface{})

	err := yaml.Unmarshal(*spec, &genericMap)
	if err != nil {
		log.Fatalf("invalid configuration file format: %v", err)
	}

	version := genericMap["version"]
	switch version {
	case "20221023":
		loadProject20221023(scope, spec)
	default:
		log.Fatalf("unsupported configuration file version: %v", version)
	}
}

func readSpecFile(dir string) *[]byte {

	configurationFile := fmt.Sprintf("%s/cevixe.yaml", dir)
	if file.Exists(configurationFile) {
		return file.GetBytes(configurationFile)
	}

	configurationFile = fmt.Sprintf("%s/cevixe.yml", dir)
	if file.Exists(configurationFile) {
		return file.GetBytes(configurationFile)
	}

	log.Fatalf("not found configuration file on project dir: %v", dir)
	return nil
}
