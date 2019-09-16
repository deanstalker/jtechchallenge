package main

import (
	"context"
	user "github.com/deanstalker/jtechchallenge/srv/user/proto"
	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/trace"
	"net/http"
	"os"
	"strings"
)

type UserREST struct {
	ctx context.Context
	wc *restful.Container
	log logrus.FieldLogger

	client user.UserService
}

type Token struct {
	UserID int64
	Username string
	jwt.StandardClaims
}

func NewUserREST(ctx context.Context, log logrus.FieldLogger, client user.UserService) *UserREST {
	return &UserREST{
		ctx: ctx,
		log: log.WithField("component", "user-rest"),
		client: client,
	}
}

func (r *UserREST) Init() *restful.WebService {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/petstore/user")
	ws.Route(ws.POST("/").To(r.create).Filter(jwtAuthentication))
	ws.Route(ws.POST("/createWithArray").To(r.batchCreate).Filter(jwtAuthentication))
	ws.Route(ws.POST("/createWithList").To(r.batchCreate).Filter(jwtAuthentication))
	ws.Route(ws.GET("/login").To(r.login))
	ws.Route(ws.GET("/logout").To(r.logout).Filter(jwtAuthentication))
	ws.Route(ws.GET("/{username}").To(r.byUsername).Filter(jwtAuthentication))
	ws.Route(ws.PUT("/{username}").To(r.update).Filter(jwtAuthentication))
	ws.Route(ws.DELETE("/{username}").To(r.delete).Filter(jwtAuthentication))
	return ws
}

func (r *UserREST) create(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "Pets.Delete")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	u := &user.UserItem{}
	if err := req.ReadEntity(u); err != nil {
		_ = rsp.WriteError(500, err)
	}

	_, err := r.client.Create(ctx, &user.CreateRequest{
		User: u,
	})
	if err != nil {
		_ = rsp.WriteError(500, err)
	}
}

func (r *UserREST) batchCreate(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "User.BatchCreate")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	users := make([]*user.UserItem, 0)
	if err := req.ReadEntity(users); err != nil {
		_ = rsp.WriteError(500, err)
	}

	_, err := r.client.BatchCreate(ctx, &user.BatchCreateRequest{
		User:users,
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
}

func (r *UserREST) login(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "User.Login")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	username := req.QueryParameter("username")
	password := req.QueryParameter("password")

	u, err := r.client.Get(ctx, &user.GetRequest{
		Username: username,
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	if u.GetUser() == nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, "A problem occurred e")
		return
	}

	// TODO Hash this ... oh hell yes hash this. Add a secret to the config and do this properly.
	if string(u.User.Password) != password {
		_ = rsp.WriteErrorString(http.StatusForbidden, "A problem occurred p")
		return
	}

	tk := &Token{UserID: u.User.Id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))

	if err = rsp.WriteEntity(map[string]string{
		"id": string(u.User.Id),
		"username": u.User.Username,
		"email": u.User.Email,
		"token": tokenString,
	}); err != nil {
		_ = rsp.WriteErrorString(http.StatusForbidden, "Access denied")
		return
	}
}

func jwtAuthentication(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain) {
	tokenHeader := req.HeaderParameter("Authorization")
	if tokenHeader == "" {
		_ = rsp.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}

	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		_ = rsp.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}

	tokenPart := splitted[1]
	tk := &Token{}
	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}

	if !token.Valid {
		_ = rsp.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}

	chain.ProcessFilter(req, rsp)
}

func (r *UserREST) logout(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "User.Logout")
	defer tr.Finish()


}

// GET http://localhost:8080/user/{username}
//
func (r *UserREST) byUsername(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "User.ByUsername")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	username := req.PathParameter("username")
	u, err := r.client.Get(ctx, &user.GetRequest{
		Username: username,
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	if u.GetUser() == nil {
		_ = rsp.WriteErrorString(http.StatusNotFound, "User not found")
		return
	}

	if err = rsp.WriteEntity(u.User); err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
}

// PUT http://localhost:8080/user/{username}
//
func (r *UserREST) update(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "User.Update")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	username := req.PathParameter("username")
	current, err := r.client.Get(ctx, &user.GetRequest{
		Username: username,
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	if current.GetUser() == nil {
		_ = rsp.WriteErrorString(http.StatusNotFound, "User not found")
		return
	}

	u := &user.UserItem{}
	if err := req.ReadEntity(u); err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	_, err = r.client.Update(ctx, &user.UpdateRequest{
		User: u,
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
}

// DELETE http://localhost:8080/user/{username}
//
func (r *UserREST) delete(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "User.Delete")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	username := req.PathParameter("username")

	_, err := r.client.Delete(ctx, &user.DeleteRequest{
		Username: username,
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
}
