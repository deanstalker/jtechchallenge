package main

import (
	"database/sql"
	"fmt"
	"os"

	storedb "github.com/deanstalker/jtechchallenge/srv/store/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro"

	store "github.com/deanstalker/jtechchallenge/srv/store/proto"
	log "github.com/sirupsen/logrus"
)

const (
	Name          = "store"
	QualifiedName = "go.micro.srv." + Name
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

	dbc, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		log.WithError(err).Warn("Unable to open handle to database")
	}

	if err = dbc.Ping(); err != nil {
		log.WithError(err).Fatal("Unable to contact database")
	}

	logger := log.New()
	repo := storedb.NewStoreRepo(logger, dbc)

	if err := store.RegisterOrderHandler(service.Server(), NewOrderHandler(log.WithField("service", QualifiedName), repo)); err != nil {
		log.WithError(err).Fatal("Service unable to accept connections")
	}

	if err := service.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
