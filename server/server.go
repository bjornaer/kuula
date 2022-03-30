package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/decode", func(c *gin.Context) {
		username := c.PostForm("username")
		c.JSON(http.StatusCreated, gin.H{
			"model_id": username + "1234",
			"message":  "model deployed succesfully",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
