package json_output

import "os"

var permissions = os.FileMode(0644)

type Output interface {
	Write()
}
