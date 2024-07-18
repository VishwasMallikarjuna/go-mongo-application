package responses

type ErrorDetail struct {
	ErrorEventId     string `json:"errorEventId"`
	ErrorDescription string `json:"errorDescription"`
}

func NewErrorDetail(requestId string, description string) *ErrorDetail {
	errorDetail := ErrorDetail{
		ErrorEventId:     requestId,
		ErrorDescription: description,
	}
	return &errorDetail
}

func NewErrorDetailResponse(code int, requestId string, description string) *ErrorDetailResponse {
	return &ErrorDetailResponse{
		Code: code,
		Body: &ErrorDetail{
			ErrorEventId:     requestId,
			ErrorDescription: description,
		},
	}
}
