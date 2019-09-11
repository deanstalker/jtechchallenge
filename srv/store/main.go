package main

import "github.com/micro/go-micro"

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.store"),
	)

	service.Init()


}
