package main

import (
	"fmt"
	"flag"
	//_ "github.com/treeptik/datamgmt/module"
	"github.com/treeptik/datamgmt/module/logging"
)

func main() {
	//Check command line arg if --loging is specified will start logging
	loggingfunction := flag.Bool("logging", false, "a bool")
	flag.Parse()
	if *loggingfunction {
		fmt.Println("Logging enabled")
		logging.Start()
	} else {
		fmt.Println("no module enabled")
	}
}
