package models

import "github.com/gin-gonic/gin"

type HTTPError struct {
	Code    int    `json:"code"`
	Path    string `json:"path"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func NewHTTPError(ctx *gin.Context, status int, message string, err error) {
	ctx.JSON(status, HTTPError{
		Code:    status,
		Path:    ctx.Request.RequestURI,
		Message: message,
		Details: err.Error(),
	})
}
