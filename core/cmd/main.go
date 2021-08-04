package main

import (
	"x-msa-auth/core"
	"x-msa-auth/helper"
	corehelper "x-msa-core/helper"

	"github.com/0LuigiCode0/logger"
)

func main() {
	conf := &helper.Config{}
	if err := corehelper.ParseConfig(helper.ConfigDir+helper.ConfigFile, conf); err != nil {
		logger.Log.Errorf("config parse invalid: %v", err)
		corehelper.Wg.Wait()
		return
	}
	srv, err := core.InitServer(conf)
	if err != nil {
		logger.Log.Errorf("server not initialized: %v", err)
		srv.Close()
		corehelper.Wg.Wait()
		return
	}
	srv.Start()

	srv.Close()
	corehelper.Wg.Wait()
}
