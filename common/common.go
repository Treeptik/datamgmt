package common

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
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

func DeleteData(container_name, elasticsearchUrl, index string) error {
	body := strings.NewReader(`
	{
		"query": {
			"bool" : {
				"should" : [
					{ "term" : { "docker.container.name" : "`+container_name+`" } },
					{ "term" : { "beat.name" : "`+container_name+`" } },
					{ "term" : { "container_name" : "`+container_name+`" } },
					{ "term" : { "host" : "`+container_name+`" } }
				],
				"minimum_should_match" : 1,
				"boost" : 1.0
			}
		}
	}`)

	req, _ := http.NewRequest("POST", elasticsearchUrl+"/"+index+"/_delete_by_query", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
			panic(err)
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err == nil {
		//fmt.Println(string(htmlData))
		fmt.Println("Successfully delete backend data")
	}
	return err
}
