package main

import (
	"endereco"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"profissional"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	profissional.Startup(server)
	endereco.Startup(server)

	err := server.Run()

	if err != nil {
		log.Fatal(err)
	}
}
