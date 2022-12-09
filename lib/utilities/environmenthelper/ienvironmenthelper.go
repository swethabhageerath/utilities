package environmenthelper

type IEnvironmentHelper interface {
	Get(key string) string
	Set(key string, value string) error
}
