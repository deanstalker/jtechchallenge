package main

import (
	"context"
	"github.com/deanstalker/jtechchallenge/srv/pet/db"
	pet "github.com/deanstalker/jtechchallenge/srv/pet/proto"
	"github.com/micro/go-micro/errors"
	"github.com/sirupsen/logrus"
)

type PetHandler struct {
	log logrus.FieldLogger

	repo *db.DefaultPetRepo
}

var _ pet.PetServiceHandler = (*PetHandler)(nil)

func NewPetHandler(log logrus.FieldLogger, repo *db.DefaultPetRepo) *PetHandler {
	return &PetHandler{
		log:  log,
		repo: repo,
	}
}

func (h PetHandler) Add(ctx context.Context, req *pet.AddRequest, rsp *pet.AddResponse) error {
	if req.GetPet() == nil {
		return errors.BadRequest("missing_pet", "A valid pet is missing from the request")
	}

	if err := h.repo.Add(req.Pet); err != nil {
		return err
	}

	return nil
}

func (h PetHandler) Update(ctx context.Context, req *pet.UpdateRequest, rsp *pet.UpdateResponse) error {
	if req.GetPet() == nil {
		return errors.BadRequest("missing_pet", "An existing pet is required before updating")
	}

	if err := h.repo.Update(req.Pet); err != nil {
		return err
	}

	return nil
}

func (h PetHandler) ByStatus(ctx context.Context, req *pet.ByStatusRequest, rsp *pet.ByStatusResponse) error {
	if req.Status == "" {
		return errors.BadRequest("missing_status", "Please enter a valid status")
	}

	pets, err := h.repo.ByStatus(req.Status)
	if err != nil {
		return err
	}

	rsp.Pets = pets

	return nil
}

func (h PetHandler) ByID(ctx context.Context, req *pet.ByIDRequest, rsp *pet.ByIDResponse) error {
	if req.Id == 0 {
		return errors.BadRequest("missing_id", "Please enter a valid pet id")
	}

	p, err := h.repo.ByID(req.Id)
	if err != nil {
		return err
	}

	rsp.Pet = p

	return nil
}

func (h PetHandler) Delete(ctx context.Context, req *pet.DeleteRequest, rsp *pet.DeleteResponse) error {
	if req.Id == 0 {
		return errors.BadRequest("missing_id", "Please enter a valid pet id")
	}

	if err := h.repo.Delete(req.Id); err != nil {
		return err
	}

	return nil
}

func (h PetHandler) UploadImage(ctx context.Context, req *pet.UploadImageRequest, rsp *pet.UploadImageResponse) error {
	return nil
}
