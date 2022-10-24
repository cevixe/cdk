package spec

type Handler struct {
	Name     string       `field:"required" yaml:"name"`
	Type     Handler_Type `field:"required" yaml:"type"`
	Events   *[]string    `field:"required" yaml:"events"`
	Commands *[]string    `field:"required" yaml:"commands"`
}

type Handler_Type string

const (
	Handler_Basic    Handler_Type = "basic"
	Handler_Standard Handler_Type = "standard"
	Handler_Advanced Handler_Type = "advanced"
)
