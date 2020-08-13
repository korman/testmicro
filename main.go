package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/korman/testmicro/handler"
	"github.com/korman/testmicro/subscriber"

	testmicro "github.com/korman/testmicro/proto/testmicro"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.testmicro"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	testmicro.RegisterTestmicroHandler(service.Server(), new(handler.Testmicro))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.testmicro", service.Server(), new(subscriber.Testmicro))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.testmicro", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
