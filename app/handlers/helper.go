package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Result(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, gin.H{"code": code, "data": data, "msg": msg})
}

func ResultCreated(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "data": "", "msg": ""})
}

func ResultOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": data, "msg": ""})
}

func ResultList(ctx *gin.Context, data interface{}, total int64) {
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "rows": data, "msg": "", "total": total})
}

func ResultOkMsg(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": data, "msg": msg})
}

func ResultFail(ctx *gin.Context, code int, err interface{}) {
	ctx.JSON(code, gin.H{"code": code, "data": nil, "msg": err})
}

func ResultFailData(ctx *gin.Context, data interface{}, err interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": data, "msg": err})
}

func ResultSlient(ctx *gin.Context) {
	ctx.String(http.StatusOK, "")
}
