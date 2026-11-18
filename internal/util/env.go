package util

import (
	"os"
	"strings"
)

// Gets the value for a given environment variable
func GetEnv(env_name string) string {
	return os.Getenv(strings.ToUpper(env_name))
}
