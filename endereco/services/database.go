package services

import (
	"context"
	"endereco/ent"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var DbClient *ent.Client

func InitDatabase() {
	client, err := ent.Open("mysql", os.Getenv("DATABASE_URL")+"/enderecos?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Fatal(err)
	}

	DbClient = client
}
