package errors

import "net/http"

type RestErr struct {
	Message string
	Status  int
	Error   string
}

func NewInternelServerError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   "Internel_Server_Error",
	}
}

func NewBadReqstError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "Bad_reqst_Error",
	}
}
