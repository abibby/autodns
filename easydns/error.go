package easydns

type Error struct {
	Body *ErrorBody `json:"error"`
}

type ErrorBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return err.Body.Message
}
