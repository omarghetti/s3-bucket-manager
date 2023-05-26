package main

import (
	"example/s3-bucket-manager/pkg/logger"
	"example/s3-bucket-manager/pkg/setting"
	routers "example/s3-bucket-manager/router"
	"log"
	"net/http"
)

func init() {
	setting.SetupEnv()
	logger.InitLogger()
}

func main() {
	initRouters := routers.InitRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: initRouters,
	}

	log.Printf("Listening on port %s", server.Addr)

	server.ListenAndServe()
}
