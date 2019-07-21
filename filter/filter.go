package filter

type filter interface {
	Filter(requestBody []byte) (bool, error)
}
