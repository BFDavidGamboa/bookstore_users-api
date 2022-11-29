package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusRequestedRangeNotSatisfiable,
		Error:   "not_found",
	}
}

func NewRequestedRangeNotSatisfiable(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusRequestedRangeNotSatisfiable,
		Error:   "requested_range_not_satisfiable",
	}
}
