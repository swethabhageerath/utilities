package helpers

type IEnvironmentHelper interface {
	Get(key string) string
	Set(key string, value string) error
}
