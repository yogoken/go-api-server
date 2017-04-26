// Package search provides json-formatted filepath, row and comment, which sorted by package based on the engine and constants
package search

import (
	"github.com/gin-gonic/gin"
	"golang_practice/constants"
	"golang_practice/engine"
)

// Search formats using the context with web framework, called gin.
// pattern and packageName are assigned by proper string
// It returns json constants http/ packagename and pattern.
func Search(context *gin.Context) {
	pattern := context.Param("pattern")
	packageName := context.Query("p")
	context.JSON(constants.HTTP_OK, engine.Search(packageName, pattern))
}
