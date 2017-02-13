package main

import (
	"github.com/treeptik/datamgmt/logging"
	"github.com/treeptik/datamgmt/common"
)

func DatamgmgInit() {
	client := common.ConnectDocker()
	common.CheckLogstash(client)
	common.CheckElasticsearch(client)
}

func main() {

	DatamgmgInit()

	logging.Start()
}
