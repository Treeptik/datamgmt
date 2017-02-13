package common

import (
  "golang.org/x/net/context"
  "fmt"
  "github.com/docker/docker/client"
  "github.com/docker/docker/api/types"
  "os"
  "github.com/docker/docker/api/types/container"
  "net/http"
)

func CheckLogstash(client *client.Client) {
  _, err := http.Get("http://logstash:9600/")
  if err != nil {
    fmt.Println("Cannot connect to http://logstash:9600/")
    os.Exit(1)
    /*containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{
  		Size:    true,
  		All:     true,
  		Since:   "container",
  		Filters: ContainerFilter([]string{"application-type=logstash", "origin=datamgmt"}),
  	})
    if err != nil {
      fmt.Println(err)
    } else {
      switch {
        case len(containers) < 1:
          fmt.Println("No logstash container")
          StartLogstash(client)
        case len(containers) > 1:
          fmt.Println("Weird behaviour")
        default:
          if containers[0].State != "running" {
            fmt.Println("Let's start logstash")
            err := client.ContainerStart(context.Background(), containers[0].ID, types.ContainerStartOptions{})
            if err != nil {
              fmt.Println(err)
            }
          } else {
            fmt.Println("Logstash already running, strange behaviour of logstash, api is reachable ?")
          }
      }
    }*/
  } else {
    fmt.Println("Successfully connected to logstash")
  }
}

func CheckElasticsearch(client *client.Client) {
  _, err := http.Get("http://elasticsearch:9200/")
  if err != nil {
    fmt.Println("Cannot connect to http://elasticsearch:9200/")
    //StartElasticsearch(client)
    os.Exit(1)
  } else {
    fmt.Println("Successfully connected to elasticsearch")
  }
}

func StartLogstash(client *client.Client) {
  labels := make(map[string]string)
  labels["origin"] = "datamgmt"
  labels["application-type"] = "logstash"
  r, err := client.ContainerCreate(context.Background(), &container.Config{Hostname: "datamgmt-logstash", Labels: labels, Image: "cloudunit/datamgmt-logstash:latest", Cmd: []string{"-f", "/opt/logstash/conf.d/", "--http.host", "0.0.0.0"}}, nil, nil, "cu-datamgmt-logstash") //&container.HostConfig{ VolumesFrom: "cu-datamgmt-manager"}
  if err != nil {
    fmt.Println("Could not create logstash container", err)
  }
  err = client.ContainerStart(context.Background(), r.ID, types.ContainerStartOptions{})
  if err != nil {
    fmt.Println("Cloud not start logstash container", err)
  }
}
