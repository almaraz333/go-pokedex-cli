package internal

import (
	"os"
)

func ExitCallback(config *Config, cache *Cache, params string) {
	os.Exit(1)
}
