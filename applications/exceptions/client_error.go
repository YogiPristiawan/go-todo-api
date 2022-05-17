package exceptions

type ClientError struct {
	Message    string
	StatusCode int
}
