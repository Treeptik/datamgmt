package main

import (
	"fmt"
	"flag"
	_ "github.com/treeptik/datamgmt/module"
	//"github.com/treeptik/datamgmt/module/logging"
	"github.com/treeptik/datamgmt/common"
)

func main() {
	loggingfunction := flag.Bool("logging", false, "a bool")
	flag.Parse()
	if *loggingfunction {
		fmt.Println("Logging enabled")
		common.Listener()
	} else {
		fmt.Println("Logging disabled")
	}
}
