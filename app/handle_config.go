package app

import "github.com/Abreu00/ts-generate/config_files"

func generateConfig() {
	config_files.NewTsconfig().Write()
	config_files.NewOrmConfig().Write()
}
