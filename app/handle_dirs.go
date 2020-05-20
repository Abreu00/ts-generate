package app

import (
	"os"
)

func generateDirs() {
	paths := []string{"dist", "src", "src/database", "src/database/migrations", "src/entity"}
	perm := os.FileMode(0755)

	for _, path := range paths {
		os.Mkdir(path, perm)
	}
}
