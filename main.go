package main

import (
	"AG-API-Server/functions"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Handle("GET", "/ping", functions.Ping)

	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
