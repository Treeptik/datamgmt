package logging

import (
  "bytes"
  "fmt"
  "strings"
  "bufio"
  "github.com/docker/docker/api/types"
  "github.com/docker/docker/api/types/container"
  "github.com/docker/docker/api/types/events"
  networktypes "github.com/docker/docker/api/types/network"
  "github.com/docker/docker/client"
  "golang.org/x/net/context"
)

func StartLogging(client *client.Client, message events.Message) {
  var buffer bytes.Buffer
  basecontainer := message.Actor.Attributes["name"]
  buffer.WriteString(basecontainer)
  buffer.WriteString("_filebeat")
  containername := buffer.String()
  inspect_result, err := client.ContainerInspect(context.Background(), containername)
  if err != nil  {
    if strings.Contains(err.Error(), "No such container:") {
      fmt.Println("Logging container doesn't exist lets create and start it")
      execid, err := client.ContainerExecCreate(context.Background(), basecontainer, types.ExecConfig{
		    Tty:          true,
        AttachStdout: true,
        Cmd:          []string{"id", "-u"},
	    })
      var gid string
      if err != nil {
        fmt.Println("Error When creating exec task")
      } else {
          output, err := client.ContainerExecAttach(context.Background(), execid.ID, types.ExecConfig{
            Tty:          true,
            AttachStdout: true,
            Cmd:          []string{"id", "-u"},
          })
          if err != nil {
            fmt.Println("Error on ExecAttach")
          } else {
            scanner := bufio.NewScanner(output.Reader)
            for scanner.Scan() {
              gid = scanner.Text()
	          }
            if err != nil {
              fmt.Println("Err while converting uidtemp to uid", err)
            }
          }
      }
      user := "10000:"+gid
      labels := make(map[string]string)
      labels["origin"] = "datamgmt"
      labels["application-type"] = "filebeat"
      network := map[string]*networktypes.EndpointSettings{
			    "datamgmt": {},
		  }
      r, err := client.ContainerCreate(context.Background(), &container.Config{Hostname: basecontainer, User: user, Labels: labels, Image: "cloudunit/datamgmt-filebeat:latest", Env: []string{"APPLICATION_TYPE="+message.Actor.Attributes["application-type"]}}, &container.HostConfig{ VolumesFrom: []string{basecontainer}}, &networktypes.NetworkingConfig{ EndpointsConfig: network}, containername)
      if err != nil {
        fmt.Println("Could not create filebeat container", err)
      } else {
        err = client.ContainerStart(context.Background(), r.ID, types.ContainerStartOptions{})
        if err != nil {
          fmt.Println("Cannot start filebeat container", err)
        }
      }
    } else {
      fmt.Println("Error on container inspect", err)
    }
  } else {
    if inspect_result.State.Status != "running" {
      fmt.Println("Lets start logging container")
      err := client.ContainerStart(context.Background(), inspect_result.ID, types.ContainerStartOptions{})
      if err != nil {
        fmt.Println("Cannot start filebeat container associated to ", basecontainer, err)
      }
    } else {
      fmt.Println("Container aleady running")
    }
  }
}
