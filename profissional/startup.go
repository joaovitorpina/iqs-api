package profissional

import (
	"github.com/gin-gonic/gin"
	"profissional/controllers"
	"profissional/services"
)

func Startup(server *gin.Engine) {
	services.InitDatabase()
	AddRoutes(server)
}

func AddRoutes(server *gin.Engine) {
	routes := server.Group("/profissionais")
	{
		routes.GET("/", controllers.ListarProfissionaisReduzido)
		routes.GET("/:id", controllers.BuscarProfissionalPorId)

		tratamentoRoutes := routes.Group("/:id/tratamentos")
		{
			tratamentoRoutes.GET("/", controllers.ListarTratamentosPorProfissional)
		}

		whatsappRoutes := routes.Group("/:id/whatsapps")
		{
			whatsappRoutes.GET("/", controllers.BuscarWhatsappsPorProfissional)
		}

		convenioRoutes := routes.Group("/:id/convenios")
		{
			convenioRoutes.GET("/", controllers.BuscarConveniosPorProfissional)
		}

		midiasRoutes := routes.Group("/:id/midias")
		{
			midiasRoutes.GET("/fotos", controllers.BuscarFotosPorProfissional)
			midiasRoutes.GET("/videos", controllers.BuscarVideosPorProfissional)
			midiasRoutes.GET("/podcasts", controllers.BuscarPodcastsPorProfissional)
		}
	}
}
