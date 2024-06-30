package utils

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

// DD prints the value, sends a JSON response, and returns to stop further execution
func DD(c *gin.Context, value interface{}) {
	// Print the value in the console
	spew.Dump(value)
	
	// Send a JSON response
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Debugging",
		"data":    value,
	})
	
	// Ensure the response is written and stop the execution
	c.Abort()
}
