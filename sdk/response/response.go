package response

type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Body       interface{} `json:"body"`
	StatusCode int         `json:"status_code"`
}

func ResponseWhenFail(statusCode int, message string) Response {
	return Response{
		StatusCode: statusCode,
		Success:    false,
		Message:    message,
		Body:       nil,
	}
}

func ResponseWhenSuccess(statusCode int, message string, body interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Success:    true,
		Message:    message,
		Body:       body,
	}
}
