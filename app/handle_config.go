package app

import "github.com/Abreu00/ts_tool/config_files"

func generateConfig() {
	var ts = config_files.NewTsconfig()
	ts.Write()
	var o = config_files.NewOrmConfig()
	o.Write()
}
