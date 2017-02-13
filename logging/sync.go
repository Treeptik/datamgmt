package logging

import (
  "fmt"
  "time"
  "bytes"
  "strings"
  "golang.org/x/net/context"
  "github.com/docker/docker/client"
  "github.com/docker/docker/api/types"
  "github.com/docker/docker/api/types/events"
  "github.com/treeptik/datamgmt/common"
)

//Function to create repetitive config sync
func ConfigSyncCron(client *client.Client) {
    ConfigSync(client)
    for range time.Tick(time.Minute *10){
      ConfigSync(client)
    }
}

//Function to synchronize config
func ConfigSync(client *client.Client) {
  fmt.Println("ConfigSync")
  containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{
    All:     false,
    Since:   "container",
    Filters: common.ContainerFilter([]string{"logging=enabled", "logging-type=file"}),
  })
  if err != nil {
    fmt.Println("Error while listing container")
  } else {
    if len(containers) > 0 {
      for _, container := range containers {
        containername := container.Names[0][1:]
        fmt.Println(containername)
        attributes := make(map[string]string)
        attributes["name"] = containername
        attributes["application-type"] = container.Labels["application-type"]
        actor := events.Actor {
          ID: container.ID,
          Attributes: attributes,
        }
        message := events.Message {
          Status: container.Status,
          ID:     container.ID,
          Type:   "container",
          Action: "start",
          Actor:  actor,
        }
        var buffer bytes.Buffer
        buffer.WriteString(containername)
        buffer.WriteString("_filebeat")
        containername = buffer.String()
        inspect_result, err := client.ContainerInspect(context.Background(), containername)
        if err != nil {
          if strings.Contains(err.Error(), "No such container:") {
            StartLogging(client, message)
          } else {
            fmt.Println("Weird error", err)
          }
        } else {
          for _, env := range inspect_result.Config.Env {
            apptype := strings.SplitAfter(env, "=")
            if apptype[0] == "APPLICATION_TYPE=" && apptype[1] != container.Labels["application-type"] {
              err := client.ContainerRemove(context.Background(), containername, types.ContainerRemoveOptions{
            		RemoveVolumes: true,
                Force:         true,
            	})
              if err != nil {
                fmt.Println("Error while deleting container", err)
              }
            }
          }
          StartLogging(client, message)
        }
    	}
    }
  }
}
