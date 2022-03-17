package http

import (
	"github.com/RogerDurdn/ConfigurationService/pkg/lib"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


var ds lib.DataSource

func Server(dataSource lib.DataSource)  {
	ds = dataSource
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
	c.JSON(http.StatusOK, ds.FetchDcData())
}

func getStatus(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"status":"get-status"})
}