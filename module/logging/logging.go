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

//var (
  //  "elasticsearchUrl" = {"http://elasticsearch:9200"},
  //  "logstashUrl" = {"http://logstash:9600"},
//)
var channels types.Listener

func init() {
  //common.AddModule()
  channels = common.InitListener(ModuleName)
  fmt.Println(channels)
  Start()
  //Start(channels)
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
        fmt.Println("Logging start")
        StartLogging(client, e)
      case e := <-channels.Stop:
        fmt.Println("Logging stop")
        StopLogging(client, e)
      case e := <-channels.Destroy:
        fmt.Println("Logging destroy")
        DestroyLogging(client, e)
    }
  }
}
