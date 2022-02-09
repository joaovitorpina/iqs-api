package services

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"materia/ent"
	"os"
)

func InitDatabase() {
	client := CreateDbClient()
	defer client.Close()

	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Fatal(err)
	}
}

func CreateDbClient() *ent.Client {
	client, err := ent.Open("mysql", os.Getenv("DATABASE_URL")+"/materias?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}

	return client
}
