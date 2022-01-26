package services

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"profissional/ent"
)

var DbClient *ent.Client

func InitDatabase() {
	client, err := ent.Open("mysql", os.Getenv("DATABASE_URL")+"/profissionais?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}

	DbClient = client
}
