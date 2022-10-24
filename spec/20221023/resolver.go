package spec

type Resolver struct {
	Name      string    `field:"required" yaml:"name"`
	Operation string    `field:"required" yaml:"operation"`
	Functions *[]string `field:"required" yaml:"functions"`
}
