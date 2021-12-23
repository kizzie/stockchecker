package main

import (
    "github.com/gin-gonic/gin"
	"github.com/kizzie/stockchecker/stockchecker"
)


func main() {    
	// very basic start a webservice
    router := gin.Default()
    router.GET("/", stockchecker.GetStock)

	// getStock()
    router.Run(":8080")
}
