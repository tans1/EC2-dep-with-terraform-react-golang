package utils

type ErrorResponse struct {
	StatusCode int
	Message    string
}

func (r ErrorResponse) Error() string {
	return r.Message
}

func DecodeError(err error) (int, string) {
	return err.(ErrorResponse).StatusCode, err.(ErrorResponse).Message
}
