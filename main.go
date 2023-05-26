package main

import (
	"example/s3-bucket-manager/pkg/setting"
	"example/s3-bucket-manager/routers"
	"log"
	"net/http"
)

func init() {
	setting.SetupEnv()
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
