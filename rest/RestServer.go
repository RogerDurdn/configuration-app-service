package rest

import (
	"github.com/RogerDurdn/ConfigurationService/pkg/lib"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)


func Server(dataSource lib.DataSource)  {
	r := gin.Default()

	ApiEndpoints(r.Group("/api"), dataSource)
	ActuatorEndpoints(r.Group("/status"), dataSource )
	log.Panic(r.Run(":3000"))
}

func AbortAndBadRequest(context *gin.Context) {
	context.JSON(400, gin.H{"msg": "Bad request"})
	context.Abort()
}

func GetQueryParamInt( context *gin.Context, key string)int  {
	return convertOrAbort(context, context.Query(key) )
}

func GetQueryParam(context *gin.Context, key string) string{
	return checkPresentOrAbort(context, context.Query(key))
}

func GetPathParamInt( context *gin.Context, key string)int  {
	return convertOrAbort(context, context.Param(key) )
}

func GetPathParam(context *gin.Context, key string) string{
	return checkPresentOrAbort(context, context.Param(key))
}

func convertOrAbort( context *gin.Context, value string)int  {
	if  value != ""{
		if  intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	AbortAndBadRequest(context)
	return 0
}

func checkPresentOrAbort(context *gin.Context, value string) string{
	if value == ""{
		AbortAndBadRequest(context)
	}
	return value
}
