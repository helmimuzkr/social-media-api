package handler

import (
	"social-media-app/features/user"
	"net/http"
	"strings"
)

type UserReponse struct {
	// ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		// ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
	}
}

func SuccessResponse(code int, message string, data ...interface{}) (int, interface{}) {
	response := map[string]interface{}{}
	if data != nil {
		response["data"] = data
	}
	if message != "" {
		response["message"] = message
	}
	return code, response
}

func ErrorResponse(message string) (int, interface{}) {
	response := map[string]interface{}{}
	code := -1

	if message != "" {
		response["message"] = message
	}

	if strings.Contains(message, "server") {
		code = http.StatusInternalServerError
	} else if strings.Contains(message, "format") {
		code = http.StatusBadRequest
	} else if strings.Contains(message, "not found") {
		code = http.StatusNotFound
	}
	return code, response
}
