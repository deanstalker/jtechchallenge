package main

import (
	"context"
	store "github.com/deanstalker/jtechchallenge/srv/store/proto"
	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/trace"
	"net/http"
	"strconv"
)

type StoreREST struct {
	ctx context.Context
	log logrus.FieldLogger

	client store.OrderService
}

func NewStoreREST(ctx context.Context, log logrus.FieldLogger, client store.OrderService) *StoreREST {
	return &StoreREST{
		ctx: ctx,
		log: log.WithField("component", "store-rest"),
		client: client,
	}
}

func (r *StoreREST) Init() *restful.WebService {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/petstore/store")
	ws.Route(ws.GET("/inventory").To(r.inventory))
	ws.Route(ws.POST("/order").To(r.placeOrder))
	ws.Route(ws.GET("/order/{orderId}").To(r.getOrder))
	ws.Route(ws.DELETE("/order/{orderId}").To(r.deleteOrder))
	return ws
}

// GET http://localhost:8080/store/inventory
//
func (r *StoreREST) inventory(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "Store.Inventory")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	result, err := r.client.Inventory(ctx, &store.InventoryRequest{})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	if result.GetInventory() == nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, "Inventory empty?")
		return
	}

	if err = rsp.WriteEntity(&result.Inventory); err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
}

// POST http://localhost:8080/store/order
//
func (r *StoreREST) placeOrder(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "Store.PlaceOrder")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	o := &store.OrderItem{}
	if err := req.ReadEntity(o); err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
	}

	_, err := r.client.Place(ctx, &store.PlaceRequest{
		Order: o,
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
}

// GET http://localhost:8080/store/order/{orderId}
//
func (r *StoreREST) getOrder(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "Store.GetOrder")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	reqId := req.PathParameter("orderId")
	orderId, err := strconv.Atoi(reqId)
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	order, err := r.client.Get(ctx, &store.GetRequest{
		Id: int64(orderId),
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	if order.GetOrder() == nil {
		_ = rsp.WriteErrorString(http.StatusNotFound, "Order not found")
	}

	if err = rsp.WriteEntity(order.Order); err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
}

// DELETE http://localhost:8080/store/order/{orderId}
//
func (r *StoreREST) deleteOrder(req *restful.Request, rsp *restful.Response) {
	tr := trace.New("api.v1", "Store.DeleteOrder")
	defer tr.Finish()

	ctx := contextTrace(r.ctx, tr)

	reqId := req.PathParameter("orderId")
	orderId, err := strconv.Atoi(reqId)
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	_, err = r.client.Delete(ctx, &store.DeleteRequest{
		Id: int64(orderId),
	})
	if err != nil {
		_ = rsp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
}