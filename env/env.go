package env

import (
	"fmt"
	"os"
)

func PORT() (string, error) { return find("PORT") }

func find(keys ...string) (string, error) {
	for _, key := range keys {
		if v := os.Getenv(key); v != "" {
			return v, nil
		}
	}

	return "", fmt.Errorf("env variable not found: %q", keys)
}
