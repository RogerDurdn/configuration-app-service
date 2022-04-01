package rest

import (
	"github.com/RogerDurdn/ConfigurationService/pkg/lib"
	"github.com/RogerDurdn/ConfigurationService/pkg/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var ds lib.DataSource

func ApiEndpoints(rg *gin.RouterGroup, dataSource lib.DataSource) {
	ds = dataSource
	rg.GET("/gp/:id", getDc)
	rg.POST("/gp", saveOrUpdateGroup)
	rg.DELETE("/gp/:id", deleteDc)
}

func getDc(context *gin.Context) {
	id := GetPathParamInt(context, "id")
	if context.IsAborted() {
		return
	}
	if group, err := ds.FetchGroupById(id); err != nil {
		context.JSON(http.StatusNotFound, err)
	} else {
		context.JSON(http.StatusOK, group)
	}
}

func saveOrUpdateGroup(context *gin.Context) {
	var group model.Group
	if err := context.ShouldBindJSON(&group); err != nil {
		log.Println(err)
		AbortAndBadRequest(context)
		return
	}
	responseCheckError(context, ds.CreateOrUpdateGroup(group), group)
}

func deleteDc(context *gin.Context) {
	id := GetPathParamInt(context, "id")
	if context.IsAborted() {
		return
	}
	err := ds.DeleteGroupById(id)
	responseCheckError(context, err, "")
}
