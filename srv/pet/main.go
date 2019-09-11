package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/deanstalker/jtechchallenge/srv/pet/proto"
	_ "github.com/go-sql-driver/mysq"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

const (
	QualifiedName = "go.micro.srv.pet"
	Name          = "pet"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	service := micro.NewService(
		micro.Name(QualifiedName),
	)

	service.Init()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s:%s/%s", "user", os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	if err := pet.RegisterPetServiceHandler(service.Server(), NewPetHandler(db, log.WithField("service", QualifiedName))); err != nil {
		log.WithError(err).Fatal("Service unable to accept connections")
	}
}