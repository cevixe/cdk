package spec

type DataSource struct {
	Name string          `field:"required" yaml:"name"`
	Type DataSource_Type `field:"required" yaml:"type"`
}

type DataSource_Type string

const (
	DataSource_Function DataSource_Type = "function"
	DataSource_Store    DataSource_Type = "store"
	DataSource_Mock     DataSource_Type = "mock"
)
