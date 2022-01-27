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
	client := services.CreateDbClient()

	routes := server.Group("/profissionais")
	{
		profissionaisController := controllers.ProfissionaisController{Client: client}
		routes.GET("/", profissionaisController.ListarProfissionaisReduzido)
		routes.GET("/:id", profissionaisController.BuscarProfissionalPorId)

		tratamentoRoutes := routes.Group("/:id/tratamentos")
		{
			tratamentosController := controllers.TratamentosController{Client: client}
			tratamentoRoutes.GET("/", tratamentosController.ListarTratamentosPorProfissional)
		}

		whatsappRoutes := routes.Group("/:id/whatsapps")
		{
			whatsappController := controllers.WhatsappController{Client: client}
			whatsappRoutes.GET("/", whatsappController.BuscarWhatsappsPorProfissional)
		}

		convenioRoutes := routes.Group("/:id/convenios")
		{
			convenioController := controllers.ConvenioController{Client: client}
			convenioRoutes.GET("/", convenioController.BuscarConveniosPorProfissional)
		}

		midiasRoutes := routes.Group("/:id/midias")
		{
			midiasController := controllers.MidiasController{Client: client}
			midiasRoutes.GET("/fotos", midiasController.BuscarFotosPorProfissional)
			midiasRoutes.GET("/videos", midiasController.BuscarVideosPorProfissional)
			midiasRoutes.GET("/podcasts", midiasController.BuscarPodcastsPorProfissional)
		}
	}
}
