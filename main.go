package main

import (
	"AG-API-Server/functions"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Client id 획득
	r.Handle("POST", "/check", functions.Check)

	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
