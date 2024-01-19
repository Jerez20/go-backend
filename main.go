package main

import (
	"net/http"
	"os"

	"Base1.com/go-backend/configs"
	"Base1.com/go-backend/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()

}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := gin.Default()

	routes.UsuarioRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	r.Run(":" + port)
}
