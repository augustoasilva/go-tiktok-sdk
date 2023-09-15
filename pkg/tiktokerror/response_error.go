package tiktokerror

import "fmt"

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	LogID   string `json:"log_id"`
}

func (r *ResponseError) ToString() string {
	return fmt.Sprintf("code: %s, message: %s, log_id: %s", r.Code, r.Message, r.LogID)
}
