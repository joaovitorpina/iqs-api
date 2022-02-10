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

	profissionaisController := controllers.ProfissionaisController{Client: client}
	areaSaudeController := controllers.AreaSaudeController{Client: client}
	tratamentosController := controllers.TratamentosController{Client: client}
	whatsappController := controllers.WhatsappController{Client: client}
	convenioController := controllers.ConvenioController{Client: client}
	midiasController := controllers.MidiasController{Client: client}

	siteRoutes := server.Group("/profissionais")
	{
		siteRoutes.GET("/", profissionaisController.ListarReduzido)

		profissionalRoutes := siteRoutes.Group("/:url_amigavel")
		{
			profissionalRoutes.GET("/", profissionaisController.BuscarPorUrlAmigavel)

			profissionalRoutes.GET("/convenios", convenioController.BuscarConveniosPorUrlAmigavelProfissional)
			profissionalRoutes.GET("/whatsapps", whatsappController.BuscarWhatsappsPorUrlAmigavelProfissional)

			midiasRoutes := profissionalRoutes.Group("/midias")
			{
				midiasRoutes.GET("/fotos", midiasController.BuscarFotosPorUrlAmigavelProfissional)
				midiasRoutes.GET("/videos", midiasController.BuscarVideosPorUrlAmigavelProfissional)
				midiasRoutes.GET("/podcasts", midiasController.BuscarPodcastsPorUrlAmigavelProfissional)
			}

			profissionalRoutes.GET("/tratamentos", tratamentosController.ListarTratamentosPorUrlAmigavelProfissional)
		}
	}

	adminRoutes := server.Group("/admin/profissionais")
	{

		profissionalRoutes := adminRoutes.Group("/:id")
		{
			profissionalRoutes.GET("/", profissionaisController.BuscarPorId)

			tratamentoRoutes := profissionalRoutes.Group("/tratamentos")
			{
				tratamentoRoutes.GET("/", tratamentosController.ListarTratamentosPorIdProfissional)
			}

			whatsappRoutes := profissionalRoutes.Group("/whatsapps")
			{
				whatsappRoutes.GET("/", whatsappController.BuscarWhatsappsPorIdProfissional)
			}

			convenioRoutes := profissionalRoutes.Group("/convenios")
			{
				convenioRoutes.GET("/", convenioController.BuscarConveniosPorIdProfissional)
			}

			midiasRoutes := profissionalRoutes.Group("/midias")
			{
				midiasRoutes.GET("/fotos", midiasController.BuscarFotosPorIdProfissional)
				midiasRoutes.GET("/videos", midiasController.BuscarVideosPorIdProfissional)
				midiasRoutes.GET("/podcasts", midiasController.BuscarPodcastsPorIdProfissional)
			}
		}
	}

	server.GET("/tipos-profissional", areaSaudeController.Listar)
}
