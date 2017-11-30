package gotools

import (
	"fmt"
	"os"
)

// MustGetEnv must get the env given
func MustGetEnv(name string) (env string, err error) {
	env = os.Getenv(name)
	if env == "" {
		err = fmt.Errorf("env [%s] not exist", name)
	}
	return
}
