package gotools

import (
	"os"
)

// IsFileExist check existance of the given file
func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
