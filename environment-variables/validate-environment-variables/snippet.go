package snippets

import (
	"github.com/kelseyhightower/envconfig"
)

// See https://godoc.org/github.com/kelseyhightower/envconfig
type Specification struct {
	Regular              string
	Array                []int
	Map                  map[string]int
	SplitWithUnderscores int `split_words:"true"`
	WithDefaultValue     int `default:"42"`
	OverriddenName       int `envconfig:"MANUAL_OVERRIDE"`
	EmbeddedStruct       EmbeddedSpecification
	Required             int `required:"true"`
}

type EmbeddedSpecification struct {
	Value int
}

// Reads and validates environment variables according to the Specification struct
func validateEnvironmentVariables(namePrefix string) (Specification, error) {
	var s Specification
	err := envconfig.Process(namePrefix, &s)
	return s, err
}
