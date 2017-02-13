package main

import (
	"github.com/treeptik/cloudunit-datamgmt/logging"
	"github.com/treeptik/cloudunit-datamgmt/common"
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
