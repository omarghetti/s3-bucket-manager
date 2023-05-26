package routers

import (
	v1 "example/s3-bucket-manager/router/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	apiv1 := r.Group("/api/v1")

	apiv1.PUT("/create-bucket", v1.CreateBucket)
	apiv1.PUT("/put-object", v1.PutObject)
	apiv1.GET("/list-objects", v1.ListObjects)
	apiv1.GET("/get-object", v1.GetObject)

	return r
}
