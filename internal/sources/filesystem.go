package sources

import (
	"path/filepath"
)

func GetFileNameWithoutExtension(filePath string) {
	// Get the base name of the file (without the directory path)
	filepath.Base(filePath)
}
