package utils

import (
	"os"
)

func GetEnv(name, fallback string) string {
	val := os.Getenv(name)
	if val == "" {
		return fallback
	}

	return val
}
