package db

import (
	"database/sql"

	store "github.com/deanstalker/jtechchallenge/srv/store/proto"
	"github.com/sirupsen/logrus"
)

type StoreRepo interface {
	Inventory() (map[string]int64, error)
	PlaceOrder(order *store.OrderItem) error
	GetOrder(id int64) error
	DeleteOrder(id int64) error
}

type DefaultStoreRepo struct {
	log logrus.FieldLogger
	db  *sql.DB
}

var _ StoreRepo = (*DefaultStoreRepo)(nil)

func NewStoreRepo(log logrus.FieldLogger, db *sql.DB) *DefaultStoreRepo {
	return &DefaultStoreRepo{
		log: log.WithField("component", "storerepo"),
		db:  db,
	}
}

func (r DefaultStoreRepo) Inventory() (map[string]int64, error) {
	return nil, nil
}

func (r DefaultStoreRepo) PlaceOrder(order *store.OrderItem) error {
	return nil
}

func (r DefaultStoreRepo) GetOrder(id int64) error {
	return nil
}

func (r DefaultStoreRepo) DeleteOrder(id int64) error {
	return nil
}
