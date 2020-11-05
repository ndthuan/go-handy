package env

import (
	"os"
	"strconv"
)

//Get reads the env var and returns fallback value if undefined
func Get(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}

//ToInt reads the env var and converts the value to int. In case the env var does not exist or error occurs in conversion, the fallback value is returned
func ToInt(key string, fallback int) int {
	value, err := strconv.Atoi(Get(key, string(rune(fallback))))
	if err != nil {
		return fallback
	}

	return value
}
