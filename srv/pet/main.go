package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.pet"),
	)

	service.Init()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s:%d/%s", "user", os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	pet.RegisterPetServiceHandler(service.Server(), &DefaultPet{
		db: db,
		log: log.WithField("service", "pet"),
	})
}