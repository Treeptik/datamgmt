package logging

import (
  "golang.org/x/net/context"
  "github.com/treeptik/cloudunit-datamgmt/common"
  "fmt"
  "io"
)

func Start() {
  client := common.ConnectDocker()

  go ConfigSync(client)

  start_messages, errs := client.Events(context.Background(), common.EventFitler([]string{"start"}, []string{"logging=enabled", "logging-type=file"}))
  stop_messages, errs := client.Events(context.Background(), common.EventFitler([]string{"die", "stop", "kill"}, []string{"logging=enabled", "logging-type=file"}))
  destroy_messages, errs := client.Events(context.Background(), common.EventFitler([]string{"destroy"}, []string{"logging=enabled", "logging-type=file"}))

  loop:
    for {
      select {
        case err := <-errs:
          if err != nil && err != io.EOF {
            fmt.Println(err)
          }
          break loop
        case e := <-start_messages:
          go StartLogging(client, e)
        case e := <-stop_messages:
          go StopLogging(client, e)
        case e := <-destroy_messages:
          go DestroyLogging(client, e)
      }
    }
}
