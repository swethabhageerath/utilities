package helpers

type IMarshallingHelper interface {
	Marshall(obj interface{}) (string, error)
}
