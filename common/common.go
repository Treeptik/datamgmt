package common

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func ContainerFilter(Labels []string) filters.Args {
	filters := filters.NewArgs()
	for _, label := range Labels {
		filters.Add("label", label)
	}
	return filters
}

func ConnectDocker() (*client.Client) {
	client, err := client.NewClient("unix:///var/run/docker.sock", "1.25", nil, nil)
  if err != nil {
		panic(err)
  } else {
		fmt.Println("Successfully connected to docker socket")
	}
	return client
}

func EventFitler(Events []string, Labels []string) (types.EventsOptions) {
	filters := filters.NewArgs()
  filters.Add("type", events.ContainerEventType)
	if len(Events) >	0 {
		for _, event := range Events {
			filters.Add("event", event)
		}
	}
	if len(Events) >	0 {
		for _, label := range Labels {
			filters.Add("label", label)
		}
	}

	expectedFiltersJSON := fmt.Sprintf(`{"type":{"%s":true}}`, events.ContainerEventType)

  options_filter := make(map[string]string)
  options_filter["filter"] = expectedFiltersJSON

  options := types.EventsOptions{
    Filters: filters,
  }
	return options
}
