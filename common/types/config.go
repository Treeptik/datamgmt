package types

import (
	"github.com/docker/docker/api/types/events"
)

type Listener struct {
	Start		<-chan events.Message
	Stop 		<-chan events.Message
	Destroy <-chan events.Message
}
