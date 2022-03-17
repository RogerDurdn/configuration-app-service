package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Server()  {

	r := gin.Default()

	apiEndpoints(r.Group("/api"))
	actuatorEndpoints(r.Group("/status"))
	log.Panic(r.Run(":3000"))
}

func apiEndpoints(rg *gin.RouterGroup)  {
	rg.GET("/dc", getDcConfiguration)
}

func actuatorEndpoints(rg *gin.RouterGroup)  {
	rg.GET("/", getStatus)
}

func getDcConfiguration(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"config":"some-config"})
}

func getStatus(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"status":"get-status"})
}