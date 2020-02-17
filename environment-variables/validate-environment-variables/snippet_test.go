package snippets

import (
	"os"
	"strings"
	"testing"
)

func TestValidateEnvironmentVariable(t *testing.T) {
	os.Setenv("ENV_CONFIG_REGULAR", "foo")
	os.Setenv("ENV_CONFIG_SPLIT_WITH_UNDERSCORES", "1")
	os.Setenv("ENV_CONFIG_MANUAL_OVERRIDE", "2")
	os.Setenv("ENV_CONFIG_ARRAY", "3,4")
	os.Setenv("ENV_CONFIG_MAP", "a:5,b:6")
	os.Setenv("ENV_CONFIG_EMBEDDEDSTRUCT_VALUE", "7")

	vars, err := validateEnvironmentVariables("ENV_CONFIG")
	if !strings.Contains(err.Error(), "required key ENV_CONFIG_REQUIRED missing value") {
		t.Errorf("Expected an error informing that ENV_CONFIG_REQUIRED is not set; got: %s", err.Error())
	}

	if vars.Regular != "foo" {
		t.Errorf("vars.Regular has value: %s; expected: text1", vars.Regular)
	}

	if vars.SplitWithUnderscores != 1 {
		t.Errorf("vars.SplitWithUnderscores has value: %d; expected: text2", vars.SplitWithUnderscores)
	}

	if vars.WithDefaultValue != 42 {
		t.Errorf("vars.WithDefaultValue has value: %d; expected: 42", vars.WithDefaultValue)
	}

	if vars.OverriddenName != 2 {
		t.Errorf("vars.OverriddenName has value: %d; expected: value 3", vars.OverriddenName)
	}

	if len(vars.Array) != 2 || vars.Array[0] != 3 || vars.Array[1] != 4 {
		t.Errorf("vars.Array has value: %+v; expected: [3, 4]", vars.Array)
	}

	if vars.Map["a"] != 5 || vars.Map["b"] != 6 {
		t.Errorf("vars.Map has values %+v; expected: [a:5 b:6]", vars.Map)
	}

	if vars.EmbeddedStruct.Value != 7 {
		t.Errorf("vars.EmbeddedStruct.Value has value: %d; expected: 7", vars.EmbeddedStruct.Value)
	}
}
