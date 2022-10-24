package spec

type File struct {
	Version string  `field:"required" yaml:"version"`
	Project Project `field:"required" yaml:"project"`
}
