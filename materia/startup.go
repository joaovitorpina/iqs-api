package materia

import (
	"github.com/gin-gonic/gin"
	"materia/controllers"
	"materia/services"
)

func Startup(server *gin.Engine) {
	services.InitDatabase()
	AddRoutes(server)
}

func AddRoutes(server *gin.Engine) {
	client := services.CreateDbClient()

	materiasController := controllers.MateriasController{Client: client}
	categoriasController := controllers.CategoriaController{Client: client}

	materiasRoutes := server.Group("/materias")
	{
		materiasRoutes.GET("", materiasController.ListarReduzido)
		materiasRoutes.GET("/:url_amigavel", materiasController.BuscarPorUrlAmigavel)
	}

	server.GET("/categorias", categoriasController.Listar)
}
