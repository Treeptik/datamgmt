package logging

import (
  "fmt"
  "time"
  "bytes"
  "strings"
  "golang.org/x/net/context"
  "github.com/docker/docker/client"
  "github.com/docker/docker/api/types/events"
)

func StopLogging(client *client.Client, message events.Message) {
  time.Sleep(2 * time.Second)
  var buffer bytes.Buffer
  basecontainer := message.Actor.Attributes["name"]
  buffer.WriteString(basecontainer)
  buffer.WriteString("_filebeat")
  containername := buffer.String()
  inspect_result, err := client.ContainerInspect(context.Background(), basecontainer)
  if err != nil  {
    if strings.Contains(err.Error(), "No such container:") {
    } else {
      fmt.Println("Error on inspect")
    }
  } else if inspect_result.State.Status == "running" {
    fmt.Println("Container have been restarted, do nothing")
  } else {
    timeout := 0 * time.Second
    fmt.Println("Container have been stopped, lets stop logging container")
    client.ContainerStop(context.Background(), containername, &timeout)
  }
}
