package common

import (
  "fmt"
  "github.com/treeptik/datamgmt/common/types"
  //"github.com/docker/docker/api/types/events"
  "golang.org/x/net/context"
)

/*func Conf(ModuleName string) types.Configuration {
  /*parameter := map[string]string{
    "elasticsearchUrl":  CheckHttp(Parameter["elasticsearchUrl"]),
    "logstashUrl": CheckHttp(Parameter["logstashUrl"]),
  }
  moduleConfig := types.Configuration{
    Name: ModuleName,
  }
  return moduleConfig
}*/

//Collect Config
var EventsListener []types.Listener

func InitListener(ModuleName string) types.Listener {
  client := ConnectDocker()
  start_messages, _ := client.Events(context.Background(), EventFitler([]string{"start"}, []string{"logging=enabled", "logging-type=file"}))
  stop_messages, _ := client.Events(context.Background(), EventFitler([]string{"die", "stop", "kill"}, []string{"logging=enabled", "logging-type=file"}))
  destroy_messages, _ := client.Events(context.Background(), EventFitler([]string{"destroy"}, []string{"logging=enabled", "logging-type=file"}))
  fmt.Println(ModuleName)
  listener := types.Listener{
    Start: start_messages,
    Stop: stop_messages,
    Destroy: destroy_messages,
  }
  EventsListener = append(EventsListener, listener)
  return listener
}

func Listener() {
  fmt.Println("toto", EventsListener)
  //module.StartModule(EventsListener)
  //listener.StartModule(EventListener)
  //for _, EventListener := range EventsListener {
    //listener.StartModule(EventListener)
  //}
}
  /*start_messages, errs := client.Events(context.Background(), EventFitler([]string{"start"}, []string{"logging=enabled", "logging-type=file"}))
  stop_messages, errs := client.Events(context.Background(), EventFitler([]string{"die", "stop", "kill"}, []string{"logging=enabled", "logging-type=file"}))
  destroy_messages, errs := client.Events(context.Background(), EventFitler([]string{"destroy"}, []string{"logging=enabled", "logging-type=file"}))
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
*/
