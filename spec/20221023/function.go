package spec

type Function struct {
	Name       string `field:"required" yaml:"name"`
	DataSource string `field:"required" yaml:"datasource"`
}
