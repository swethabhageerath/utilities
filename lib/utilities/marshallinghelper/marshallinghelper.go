package marshallinghelper

import "encoding/json"

type MarshallingHelper struct{}

func (m MarshallingHelper) Marshall(obj interface{}) (string, error) {
	j, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(j), nil
}
