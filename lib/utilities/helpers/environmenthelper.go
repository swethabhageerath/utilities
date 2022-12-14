package helpers

import "os"

type EnvironmentHelper struct{}

func (e EnvironmentHelper) Get(key string) string {
	return os.Getenv(key)
}

func (e EnvironmentHelper) Set(key string, value string) error {
	return os.Setenv(key, value)
}
