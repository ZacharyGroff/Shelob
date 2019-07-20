package parser

type RequestParser interface {
	Parse(requestBody []byte) error
}
