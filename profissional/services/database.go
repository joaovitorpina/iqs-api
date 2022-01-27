package services

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"profissional/ent"
)

func InitDatabase() {
	client := CreateDbClient()
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func CreateDbClient() *ent.Client {
	client, err := ent.Open("mysql", os.Getenv("DATABASE_URL")+"/profissionais?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}

	return client
}
