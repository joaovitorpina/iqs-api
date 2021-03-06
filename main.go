package main

import (
	"endereco"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "iqs-api/docs"
	"log"
	"materia"
	"profissional"
	"unidade"
)

// @title        IQS API
// @version      1.0
// @description  Essa é a API para o site do Informacão que Salva e o admin.

// @contact.name   João Vitor Goncalves Pina
// @contact.email  joaovitorpina@icloud.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	server := gin.Default()
	server.Use(cors.AllowAll())

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	profissional.Startup(server)
	endereco.Startup(server)
	unidade.Startup(server)
	materia.Startup(server)

	err := server.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	err := godotenv.Load()

	if err != nil && gin.IsDebugging() {
		log.Fatal(err)
	}
}
