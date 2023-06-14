package superpath

import (
	"os"
	"fmt"
	"path/filepath"
)

var currentDirectory string

func init() {
	absDir, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	currentDirectory = absDir
}

func GetCurrentDirectory() string {
	return currentDirectory
}


func ChangeDirectory(dir string) error {
	if !filepath.IsAbs(dir) {
		dir = filepath.Join(currentDirectory, dir)
	}

	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			dir, _ := filepath.Rel(currentDirectory, dir)
			return fmt.Errorf("%s: Directory does not exist", dir)
		}
		return err
	}

	absDir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	cleanedDir := filepath.Clean(absDir)
	currentDirectory = cleanedDir

	return nil
}
