package logging

import (
  "fmt"
  //"io"
  "github.com/treeptik/datamgmt/common"
  "github.com/treeptik/datamgmt/common/types"
  //"golang.org/x/net/context"
  //"reflect"
  //"github.com/docker/docker/api/types/events"
)

const ModuleName = "logging"

//var channels types.Listener
var (
  channels types.Listener
  parameters = map[string]string{
    "elasticsearchUrl": "http://elasticsearch:9200",
    "logstashUrl":      "http://logstash:9600",
  }
)

func init() {
  //common.AddModule()
  for _, parameter := range parameters {
    common.CheckHttp(parameter)
  }
  channels = common.InitListener(ModuleName)
  fmt.Println(channels)
  //Start()
}

//func Start(channels types.Listener) {
func Start() {
  client := common.ConnectDocker()

  CheckImages(client)

  go ConfigSync(client)

  /*start_messages, errs := client.Events(context.Background(), common.EventFitler([]string{"start"}, []string{"logging=enabled", "logging-type=file"}))
  stop_messages, errs := client.Events(context.Background(), common.EventFitler([]string{"die", "stop", "kill"}, []string{"logging=enabled", "logging-type=file"}))
  destroy_messages, errs := client.Events(context.Background(), common.EventFitler([]string{"destroy"}, []string{"logging=enabled", "logging-type=file"}))
  fmt.Println(reflect.TypeOf(start_messages))
  */
  for {
    select {
      //case err := <-errs:
      //  if err != nil && err != io.EOF {
      //    fmt.Println(err)
      //  }
      //  break loop
      case e := <-channels.Start:
        go StartLogging(client, e)
      case e := <-channels.Stop:
        go StopLogging(client, e)
      case e := <-channels.Destroy:
        go DestroyLogging(client, e)
    }
  }
}
