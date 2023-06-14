package commands

import (
	"os"
	"fmt"
	"GoShell/errors"
	"GoShell/superpath"
	"path/filepath"
)

func Touch(files []string) {
	if len(files) == 0 {
		errors.PrintError("touch", fmt.Errorf("No params specified"))
		return
	}

	for _, file := range files {
		var absName = file
		if !filepath.IsAbs(file) {
			absName = filepath.Join(superpath.GetCurrentDirectory(), file)
		}

		f, err := os.Create(absName)
		if err != nil {
			errors.PrintError("touch", fmt.Errorf("%s: Error creating file", file))
			continue
		}
		
		defer f.Close()
	}
}