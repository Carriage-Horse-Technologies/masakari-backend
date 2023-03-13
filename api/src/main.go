package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"*",
		},
		MaxAge: 24 * time.Hour,
	}))

	router.GET("/ws/:roomId", func(c *gin.Context) {
		// roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, "last-hack")
	})
	router.GET("/status", func(c *gin.Context) {
		handleGetServerStatus(c.Writer, c.Request)
	})
	router.GET("/test/bench", func(c *gin.Context) {
		Benchmark(1.0)
	})

	go h.run()
	ProviderJob()
	router.Run(":80")
}
