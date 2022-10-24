package spec

type Project struct {
	Kind        Kind       `field:"required" yaml:"kind"`
	Name        string     `field:"required" yaml:"name"`
	Description string     `field:"required" yaml:"description"`
	Domains     *[]string  `field:"optional" yaml:"domains"`
	App         *App       `field:"optional" yaml:"app"`
	Api         *Api       `field:"optional" yaml:"api"`
	Handlers    *[]Handler `field:"optional" yaml:"handlers"`
}
