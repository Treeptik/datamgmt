package logging

import (
  "bytes"
  "fmt"
  "github.com/docker/docker/api/types"
  "github.com/docker/docker/api/types/events"
  "github.com/docker/docker/client"
  "golang.org/x/net/context"
)

func DestroyLogging(client *client.Client, message events.Message) {
  var buffer bytes.Buffer
  basecontainer := message.Actor.Attributes["name"]
  buffer.WriteString(basecontainer)
  buffer.WriteString("_filebeat")
  containername := buffer.String()
  err := client.ContainerRemove(context.Background(), containername, types.ContainerRemoveOptions{
		RemoveVolumes: true,
    Force:         true,
	})
  if err != nil {
    fmt.Println("Error while deleting container", err)
  } else {
    fmt.Println("Remove logging container")
  }
}
