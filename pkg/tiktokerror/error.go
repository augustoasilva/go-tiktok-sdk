package tiktokerror

import "fmt"

type Error struct {
	StatusCode    int    `json:"status_code"`
	Title         string `json:"title"`
	Detail        string `json:"detail"`
	OriginalError error
}

func New(statusCode int, title, detail string, originalErr error) *Error {
	return &Error{
		StatusCode:    statusCode,
		Title:         title,
		Detail:        detail,
		OriginalError: originalErr,
	}
}

func NewWithoutOriginalError(statusCode int, title, detail string) *Error {
	return New(statusCode, title, detail, nil)
}

func (e *Error) Error() string {
	msg := fmt.Sprintf("status code: %d. %s: %s.", e.StatusCode, e.Title, e.Detail)
	if e.OriginalError != nil {
		return fmt.Sprintf("%s originally caused by: %s", msg, e.OriginalError.Error())
	}
	return msg
}
