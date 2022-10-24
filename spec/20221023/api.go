package spec

type Api struct {
	DataSources *[]Resolver `field:"optional" yaml:"datasources"`
	Functions   *[]Resolver `field:"optional" yaml:"functions"`
	Resolvers   *[]Resolver `field:"optional" yaml:"resolvers"`
}
