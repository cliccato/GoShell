package commands

import (
	"os"
	"fmt"
	"GoShell/errors"
	"GoShell/superpath"
	"path/filepath"
)

func Mkdir(directories []string) {
	if len(directories) == 0 {
		errors.PrintError("touch", fmt.Errorf("No params specified"))
		return
	}

	for _, dir := range directories {
		var absName = dir
		if !filepath.IsAbs(dir) {
			absName = filepath.Join(superpath.GetCurrentDirectory(), dir)
		}

		err := os.Mkdir(absName, 0755)
		if err != nil {
			errors.PrintError("mkdir", fmt.Errorf("%s: Error creating directory", dir))
			continue
		}
	}
}