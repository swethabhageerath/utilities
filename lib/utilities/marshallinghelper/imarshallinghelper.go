package marshallinghelper

type IMarshallingHelper interface {
	Marshall(obj interface{}) (string, error)
}
