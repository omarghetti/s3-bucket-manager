package routers

import (
	v1 "example/s3-bucket-manager/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	apiv1 := r.Group("/api/v1")

	apiv1.POST("/create-bucket", v1.CreateBucket)
	apiv1.POST("/put-object", v1.PutObject)
	apiv1.GET("/list-objects", v1.ListObjects)

	return r
}
