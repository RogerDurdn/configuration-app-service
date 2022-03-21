package rest

import (
	"github.com/RogerDurdn/ConfigurationService/pkg/lib"
	"github.com/RogerDurdn/ConfigurationService/pkg/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var ds lib.DataSource

func ApiEndpoints(rg *gin.RouterGroup, dataSource lib.DataSource)  {
	ds = dataSource
	rg.GET("/gp/:id", getDc)
	rg.POST("/gp", saveOrUpdateGroup)
	rg.DELETE("/gp/:id", deleteDc)
	rg.GET("/gp/:id/feature/*other", getDcFeature)
	rg.POST("/gp/:id/feature", createOrUpdateFeature)
	rg.DELETE("/gp/:id/feature/*other", deleteDcFeature)
}

func getDc(context *gin.Context)  {
	id := GetPathParamInt(context, "id")
	if context.IsAborted(){
		return
	}
	context.JSON(http.StatusOK, ds.FetchGroupById(id))
}

func saveOrUpdateGroup(context *gin.Context) {
	var group model.Group
	if err := context.ShouldBindJSON(&group); err != nil{
		log.Println(err)
		AbortAndBadRequest(context)
		return
	}
	context.JSON(http.StatusOK, ds.CreateOrUpdateGroup(group))
}


func deleteDc(context *gin.Context) {
	id := GetPathParamInt(context, "id")
	if context.IsAborted(){
		return
	}
	ds.DeleteGroupById(id)
	context.JSON(http.StatusOK,"")
}


func getDcFeature(context *gin.Context) {
	id := GetPathParamInt(context, "id")
	key := GetQueryParam(context, "key")
	if context.IsAborted(){
		return
	}
	context.JSON(http.StatusOK, ds.FetchGroupFeatureByKey(id, key))
}

func createOrUpdateFeature(context *gin.Context) {
	var feature model.Feature
	if err := context.ShouldBindJSON(&feature); err != nil{
		log.Println(err)
		AbortAndBadRequest(context)
		return
	}
	context.JSON(http.StatusOK, ds.CreateOrUpdateFeature(feature))
}


func deleteDcFeature(context *gin.Context) {
	id := GetPathParamInt(context, "id")
	key := GetQueryParam(context, "key")
	if context.IsAborted(){
		return
	}
	ds.DeleteFeatureByKey(id, key)
	context.JSON(http.StatusOK, "")
}