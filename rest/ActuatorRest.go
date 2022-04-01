package rest

import (
	"github.com/RogerDurdn/ConfigurationService/pkg/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ActuatorEndpoints(rg *gin.RouterGroup, dataSource lib.DataSource) {
	rg.GET("/", getStatus)
}

func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "get-status"})
}
