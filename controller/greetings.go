package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greetings(c *gin.Context) {
	// renderer.Render(c, posts, http.StatusOK)
	c.JSON(http.StatusOK, gin.H{"msg": "hello world"})
}
