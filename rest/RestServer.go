package rest

import (
	"github.com/RogerDurdn/ConfigurationService/pkg/lib"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Server(dataSource lib.DataSource) {
	r := gin.Default()
	ApiEndpoints(r.Group("/api"), dataSource)
	ActuatorEndpoints(r.Group("/status"), dataSource)
	log.Panic(r.Run(":9093"))
}

func AbortAndBadRequest(context *gin.Context) {
	context.JSON(400, gin.H{"msg": "Bad request"})
	context.Abort()
}

func GetPathParamInt(context *gin.Context, key string) int {
	return convertOrAbort(context, context.Param(key))
}

func convertOrAbort(context *gin.Context, value string) int {
	if value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	AbortAndBadRequest(context)
	return 0
}

func responseCheckError(context *gin.Context, err error, body interface{}) {
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"msg": err.Error()})
	} else {
		context.JSON(http.StatusOK, body)
	}
}
