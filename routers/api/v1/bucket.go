package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBucket(ctx *gin.Context) {
	ctx.String(http.StatusOK, "CreateBucket")
}

func PutObject(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PutObject")
}

func ListObjects(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ListObjects")
}
