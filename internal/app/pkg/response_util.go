package pkg

import (
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/domain/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any, D any](status constant.ResponseStatus, errMsg T, data D) dto.APIResponse[T, D] {
	return BuildResponse_(status.GetResponseStatus(), status.GetResponseMessage(), errMsg, data)
}

func BuildResponse_[T any, D any](status string, message string, errMsg T, data D) dto.APIResponse[T, D] {
	return dto.APIResponse[T, D]{
		Status:       status,
		Message:      message,
		ErrorMessage: errMsg,
		Data:         data,
	}
}
