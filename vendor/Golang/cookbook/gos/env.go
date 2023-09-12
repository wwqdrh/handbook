package gos

import "os"

func GetEnv(env string) string {
	return os.Getenv(env)
}
