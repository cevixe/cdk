package project

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cevixe/cdk/common/file"
	"gopkg.in/yaml.v3"
)

type GenericFile struct {
	Version string `field:"required" yaml:"version"`
}

func Load(scope constructs.Construct) {

	dir := os.Getenv("CEVIXE_WORKSPACE")
	spec := readSpecFile(dir)
	file := &GenericFile{}

	err := yaml.Unmarshal(*spec, file)
	if err != nil {
		log.Fatalf("invalid configuration file format: %v", err)
	}

	switch file.Version {
	case "20221023":
		loadProject20221023(scope, spec)
	default:
		log.Fatalf("unsupported configuration file version: %v", file.Version)
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
