package marshallinghelper

type IMarshallingHelper interface {
	MarshallJson(obj interface{}) string
}
