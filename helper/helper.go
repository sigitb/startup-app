package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code`
	Status  string `json:"status"`
}

func ApiRespone(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatterValidationError(err error) []string {
	var errors []string
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			errors = append(errors, e.Error())
		}
		return errors

}

func FAILED(c *gin.Context, message string) {
	
	c.JSON(http.StatusBadRequest, struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  false,
		Message: message,
		Data:    nil,
	})
	return
}