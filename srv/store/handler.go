package main

import (
	"context"

	"github.com/deanstalker/jtechchallenge/srv/store/db"

	store "github.com/deanstalker/jtechchallenge/srv/store/proto"
	"github.com/sirupsen/logrus"
)

type OrderHandler struct {
	log  logrus.FieldLogger
	repo db.StoreRepo
}

var _ store.OrderHandler = (*OrderHandler)(nil)

func NewOrderHandler(log logrus.FieldLogger, repo db.StoreRepo) *OrderHandler {
	return &OrderHandler{
		log:  log,
		repo: repo,
	}
}

func (h *OrderHandler) Inventory(context.Context, *store.InventoryRequest, *store.InventoryResponse) error {
	return nil
}

func (h *OrderHandler) Place(context.Context, *store.PlaceRequest, *store.PlaceResponse) error {
	return nil
}

func (h *OrderHandler) Get(context.Context, *store.GetRequest, *store.GetResponse) error {
	return nil
}

func (h *OrderHandler) Delete(context.Context, *store.DeleteRequest, *store.DeleteResponse) error {
	return nil
}
