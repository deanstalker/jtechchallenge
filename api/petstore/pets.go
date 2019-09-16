package main

import (
	"context"
	srvpet "github.com/deanstalker/jtechchallenge/srv/pet/proto"
	"github.com/emicklei/go-restful"
	"github.com/gorilla/schema"
	"github.com/micro/go-micro/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/trace"
	"net/http"
	"strconv"
)

var decoder *schema.Decoder

type PetREST struct {
	ctx context.Context
	log logrus.FieldLogger

	client srvpet.PetService
}

func NewPetREST(ctx context.Context, log logrus.FieldLogger, client srvpet.PetService) *PetREST {
	return &PetREST{
		ctx: ctx,
		log: log.WithField("component", "pet-rest"),
		client: client,
	}
}

func (r *PetREST) Init() *restful.WebService {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/petstore/pet")
	ws.Route(ws.POST("").To(r.add).
		Doc("Create user").
		Reads(srvpet.Pet{}).
		Returns(200, "", srvpet.Pet{}))
	ws.Route(ws.PUT("").To(r.update))
	ws.Route(ws.GET("/findByStatus").To(r.byStatus))
	ws.Route(ws.GET("/{petId}").To(r.byID))
	ws.Route(ws.POST("/{petId}").Consumes("application/x-www-form-urlencoded").To(r.updateFromForm))
	ws.Route(ws.DELETE("/{petId}").To(r.delete))
	ws.Route(ws.POST("/{petId}/uploadImage").To(r.uploadImage))
	return ws
}


// POST http://localhost:8080/user
//
func (r *PetREST) add(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "Pets.Add")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	pet := &srvpet.Pet{}

	if err := req.ReadEntity(pet); err != nil {
		_ = rsp.WriteErrorString(http.StatusMethodNotAllowed, "Unable to read payload")
	}

	_, err := r.client.Add(ctx, &srvpet.AddRequest{
		Pet: pet,
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
}

// PUT http://localhost:8080/user
//
func (r *PetREST) update(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "Pets.Update")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	updated := &srvpet.Pet{}
	if err := req.ReadEntity(updated); err != nil {
		_ = rsp.WriteError(http.StatusMethodNotAllowed, err)
	}

	result, err := r.client.ByID(ctx, &srvpet.ByIDRequest{
		Id: updated.Id,
	})
	if err != nil {
		_ = rsp.WriteError(http.StatusBadRequest, err)
		return
	}

	if result.GetPet() == nil {
		_ = rsp.WriteError(http.StatusNotFound, errors.NotFound("not_found", "Pet not found"))
		return
	}

	_, err = r.client.Update(ctx, &srvpet.UpdateRequest{
		Pet: updated,
	})
	if err != nil {
		_ = rsp.WriteError(http.StatusBadRequest, err)
		return
	}
}

// POST http://localhost:8080/pet/{petId}
// @TODO Review this sucker, may not be 100% right
func (r *PetREST) updateFromForm(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "Pets.UpdateFromForm")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	// Get the original pet record (if it exists)
	reqId := req.PathParameter("petId")
	petId, err := strconv.Atoi(reqId)
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
	}

	currentPet, err := r.client.ByID(ctx, &srvpet.ByIDRequest{
		Id: int64(petId),
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	if currentPet.GetPet() == nil {
		_ = rsp.WriteErrorString(http.StatusNotFound, "Pet not found")
		return
	}

	if err := req.Request.ParseForm(); err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	// decode the post form data
	p := new(srvpet.Pet)
	if err := decoder.Decode(p, req.Request.PostForm); err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	_, err = r.client.Update(ctx, &srvpet.UpdateRequest{
		Pet: p,
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
	}
}

// GET http://localhost:8080/pet/findByStatus
//
func (r *PetREST) byStatus(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "Pets.ByStatus")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	status := req.QueryParameter("status")

	pet, err := r.client.ByStatus(ctx, &srvpet.ByStatusRequest{
		Status: status,
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	if err = rsp.WriteEntity(pet); err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
}

// GET http://localhost:8080/pet/{petId}
//
func (r *PetREST) byID(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "Pets.ByID")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	reqId := req.PathParameter("petId")
	petId, err := strconv.Atoi(reqId)
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	pet, err := r.client.ByID(ctx, &srvpet.ByIDRequest{
		Id: int64(petId),
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	if err = rsp.WriteEntity(pet); err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
}

func (r *PetREST) delete(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "Pets.Delete")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	reqId := req.PathParameter("petId")
	petId, err := strconv.Atoi(reqId)
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	_, err = r.client.Delete(ctx, &srvpet.DeleteRequest{
		Id: int64(petId),
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
}

func (r *PetREST) uploadImage(req *restful.Request, rsp *restful.Response) {
	// TODO Image processing isn't in scope of this challenge, i'm sure. But in essence I would do the following:
	// - I'd accept the image as a stream and capture the original filename
	// - I would then stream the image to S3, and capture the URL returned
	// - Then the photo record would be stored away against the Pet.
}


