package main

import (
	"context"
	pet "github.com/deanstalker/jtechchallenge/srv/pet/proto"
	store "github.com/deanstalker/jtechchallenge/srv/store/proto"
	user "github.com/deanstalker/jtechchallenge/srv/user/proto"
	"github.com/emicklei/go-restful"
	"github.com/google/uuid"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/web"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/trace"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.api.petstore"),
	)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	storeClient := store.NewOrderService("go.micro.srv.store", client.DefaultClient)
	petClient := pet.NewPetService("go.micro.srv.pet", client.DefaultClient)
	userClient := user.NewUserService("go.micro.srv.user", client.DefaultClient)

	logger := log.New()
	wc := restful.NewContainer()
	wc.Add(NewPetREST(ctx, logger, petClient).Init())
	wc.Add(NewStoreREST(ctx, logger, storeClient).Init())
	wc.Add(NewUserREST(ctx, logger, userClient).Init())

	service.Handle("/", wc)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}


func contextTrace(ctx context.Context, tr trace.Trace) context.Context {
	tctx := trace.NewContext(ctx, tr)

	md, ok := metadata.FromContext(tctx)
	if !ok {
		md = metadata.Metadata{}
	}

	traceID := uuid.New()
	tmd := metadata.Metadata{}
	for k,v := range md {
		tmd[k] = v
	}

	tmd["traceID"] = traceID.String()
	tmd["fromName"] = "api.v1"
	return metadata.NewContext(tctx, tmd)
}