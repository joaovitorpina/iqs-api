package endereco

import (
	"endereco/controllers"
	"endereco/services"
	"github.com/gin-gonic/gin"
)

func Startup(server *gin.Engine) {
	services.InitDatabase()
	AddRoutes(server)
}

func AddRoutes(server *gin.Engine) {
	routes := server.Group("/enderecos")
	{
		routes.GET("/:id", controllers.BuscarEnderecoPorId)
	}
}
