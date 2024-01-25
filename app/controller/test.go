package controller

import (
	"go-project/app/lib/response"
	"go-project/app/openapi"

	"github.com/gin-gonic/gin"
)

func SayHello(ctx *gin.Context) {
	resp := response.Gin{Ctx: ctx}

	commonResponse := openapi.CommonResponse{
		Status:  openapi.RESPONSESTATUS_SUCCESS,
		Message: "holle~",
	}

	resp.Success(commonResponse)
}
