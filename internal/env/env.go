package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Load .env once when the package is imported. It is safe to ignore the error
// so the app still runs in environments where a .env file is not present.
func init() {
	_ = godotenv.Load()
}

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return i
}
