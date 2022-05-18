package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status  int    `json:"rescode"`
		Message string `json:"message"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(ctx echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "success"
	response.Data = param
	return ctx.JSON(http.StatusOK, response)
}
func NewErrorResponse(ctx echo.Context, status int) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "something not right"
	return ctx.JSON(status, response)
}
