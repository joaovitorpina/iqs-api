package unidade

import (
	"github.com/gin-gonic/gin"
	"unidade/services"
)

func Startup(server *gin.Engine) {
	services.InitDatabase()
	addRoutes(server)
}

func addRoutes(server *gin.Engine) {
	client := services.CreateDbClient()

	routes := server.Group("/unidades")
	{

	}
}
