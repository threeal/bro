package utils

import "os"

func GetEnvOrDefault(key string, def string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}
	return def
}
