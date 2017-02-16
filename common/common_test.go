package common

import (
	"testing"
	"fmt"
)

func TestEventFitler(t *testing.T) {
	cases := []struct {
		event    []string
		labels   []string
	}{
		{
			event:		[]string{"start"},
			labels:		[]string{"logging=enabled", "logging-type=file"},
		},
		{
			event:  []string{"die", "stop", "kill"},
			labels:	[]string{"logging=enabled", "logging-type=file"},
		},
		{
			event:	[]string{"destroy"},
			labels:	[]string{"logging=enabled", "logging-type"},
		},
	}
	for _, c := range cases {
		option := EventFitler(c.event, c.labels)
		fmt.Println(option.Filters)
	}
}
