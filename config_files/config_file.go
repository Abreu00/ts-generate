package config_files

import "os"

var permissions = os.FileMode(0644)

type CongigFile interface {
	Write()
}
