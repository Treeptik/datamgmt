package types

import (
	"github.com/docker/docker/api/types/events"
)

type Configuration struct {}

type ModuleStart interface {
	Start()
}

type Listener struct {
	Start		<-chan events.Message
	Stop 		<-chan events.Message
	Destroy <-chan events.Message
	Function	StartJob
}

type StartJob interface {
	Start() ()
}
