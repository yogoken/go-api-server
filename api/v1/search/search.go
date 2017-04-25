package search

import (
	"github.com/gin-gonic/gin"
	"golang_practice/engine"
	"golang_practice/constants"
)

func Search(context *gin.Context) {
	pattern := context.Param("pattern")
	packageName := context.Query("p")
	context.JSON(constants.HTTP_OK, engine.Search(packageName, pattern))
}