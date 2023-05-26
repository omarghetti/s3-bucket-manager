package v1

import (
	"example/s3-bucket-manager/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBucketRequest struct {
	BucketName string `json:"bucket"`
}

type PutObjectRequest struct {
	BucketName string `json:"bucket"`
	Key        string `json:"key"`
	Body       string `json:"body"`
}

func CreateBucket(ctx *gin.Context) {
	var request CreateBucketRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Logger.Error(err.Error())
	}
	var bucketName = bucket_service.CreateBucket(request.BucketName)
	ctx.String(http.StatusOK, "Bucket Created: %s", bucketName)
}

func PutObject(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PutObject")
}

func ListObjects(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ListObjects")
}

func GetObject(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetObject")
}
