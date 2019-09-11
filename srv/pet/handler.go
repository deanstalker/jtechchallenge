package main

import (
	"context"
	"database/sql"
	pet "github.com/deanstalker/jtechchallenge/srv/pet/proto"
	"github.com/sirupsen/logrus"
)

type PetHandler struct {
	db *sql.DB
	log logrus.FieldLogger
}


var _ pet.PetServiceHandler = (*PetHandler)(nil)

func NewPetHandler(db *sql.DB, log logrus.FieldLogger) *PetHandler {
	return &PetHandler{
		db: db,
		log: log,
	}
}

func (h PetHandler) Add(ctx context.Context, req *pet.AddRequest, rsp *pet.AddResponse) error {
	return nil
}

func (h PetHandler) Update(ctx context.Context, req *pet.UpdateRequest, rsp *pet.UpdateResponse) error {
	return nil
}

func (h PetHandler) ByStatus(ctx context.Context, req *pet.ByStatusRequest, rsp *pet.ByStatusResponse) error {
	return nil
}

func (h PetHandler) ByID(ctx context.Context, req *pet.ByIDRequest, rsp *pet.ByIDResponse) error {
	return nil
}

func (h PetHandler) Delete(ctx context.Context, req *pet.DeleteRequest, rsp *pet.DeleteResponse) error {
	return nil
}

func (h PetHandler) UploadImage(ctx context.Context, req *pet.UploadImageRequest, rsp *pet.UploadImageResponse) error {
	return nil
}