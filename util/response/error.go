package response

type customError struct {
	msg string
}

func (c *customError) Error() string {
	return c.msg
}

func Custom(s string) *customError {
	return &customError{s}
}
