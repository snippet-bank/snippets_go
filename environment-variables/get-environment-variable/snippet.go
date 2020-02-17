package snippets

import (
	"os"
)

func getEnvVariable(name string) string {
	return os.Getenv(name)
}
