package main

import (
	"context"

	"github.com/deanstalker/jtechchallenge/srv/user/db"
	user "github.com/deanstalker/jtechchallenge/srv/user/proto"
	"github.com/micro/go-micro/errors"
	"github.com/sirupsen/logrus"
)

//UserHandler base user handler struct
type UserHandler struct {
	log  logrus.FieldLogger
	repo *db.DefaultUserRepo
}

var _ user.UserHandler = (*UserHandler)(nil)

//NewUserHandler handles RPC/Json API requests for users
func NewUserHandler(log logrus.FieldLogger, repo *db.DefaultUserRepo) *UserHandler {
	return &UserHandler{
		log:  log,
		repo: repo,
	}
}

//Create user
func (h *UserHandler) Create(ctx context.Context, req *user.CreateRequest, rsp *user.CreateResponse) error {
	if req.GetUser() == nil {
		return errors.BadRequest("missing_user", "Unable to create user, the request is invalid")
	}

	total, err := h.repo.Create(req.User)
	if err != nil {
		return err
	}

	if total == 1 {
		rsp.Success = true
	}

	return nil
}

//BatchCreate users
func (h *UserHandler) BatchCreate(ctx context.Context, req *user.BatchCreateRequest, rsp *user.BatchCreateResponse) error {
	if req.GetUsers() == nil {
		return errors.BadRequest("empty_users", "Users not provided")
	}

	created, err := h.repo.BatchCreate(req.Users)
	if err != nil {
		return errors.BadRequest("users_not_created", "")
	}

	if int(created) != len(req.Users) {
		rsp.Success = false
	}

	rsp.Success = true

	return nil
}

//Get user by username
func (h *UserHandler) Get(ctx context.Context, req *user.GetRequest, rsp *user.GetResponse) error {
	u, err := h.repo.ByUsername(req.Username)
	if err != nil {
		return err
	}

	rsp.User = u

	return nil
}

//Update user
func (h *UserHandler) Update(ctx context.Context, req *user.UpdateRequest, rsp *user.UpdateResponse) error {
	if req.GetUser() == nil {
		return errors.BadRequest("empty_user", "User not found")
	}

	u, err := h.repo.ByUsername(req.User.Username)
	if err != nil {
		return err
	}

	if u == nil {
		return errors.BadRequest("empty_user", "User not found")
	}

	if err := h.repo.Update(req.User); err != nil {
		return err
	}

	return nil
}

//Delete user
func (h *UserHandler) Delete(ctx context.Context, req *user.DeleteRequest, rsp *user.DeleteResponse) error {
	if req.Username == "" {
		return errors.BadRequest("empty_user", "User not found")
	}

	u, err := h.repo.ByUsername(req.Username)
	if err != nil {
		return err
	}

	if u == nil {
		return errors.BadRequest("empty_user", "User not found")
	}

	if err := h.repo.Delete(req.Username); err != nil {
		return err
	}

	return nil
}
