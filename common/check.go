package common

import (
  "fmt"
  "net/http"
  "os"
)

func CheckHttp(url string) string {
  _, err := http.Get(url)
  if err != nil {
    fmt.Println("Cannot connect to ", url)
    os.Exit(1)
  } else {
    fmt.Println("Successfully connected to ", url)
  }
  return url
}
