package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	testmicro "github.com/korman/testmicro/proto/testmicro"
)

type Testmicro struct{}

func (e *Testmicro) Handle(ctx context.Context, msg *testmicro.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *testmicro.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
