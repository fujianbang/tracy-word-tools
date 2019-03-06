package main

import (
	"check/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodPost, http.MethodOptions, http.MethodGet},
		AllowHeaders:     []string{"Origin", "Content-Type"},
	}))

	r.POST("/upload", handler.DownloadHandler)
	r.Run(":9090")
}
