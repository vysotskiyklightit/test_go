package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadJSONToModel(c *gin.Context, model ModelInterface) {
	if err := c.ShouldBindJSON(model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
