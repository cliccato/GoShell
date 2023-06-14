package commands

import (
	"GoShell/superpath"
	"GoShell/errors"
	"os"
	"fmt"
	"path/filepath"
)

var isRecursive = false
var isForced = false

func Rm(names []string, options []string) {
	if len(names) == 0 {
		errors.PrintError("rm", fmt.Errorf("No params specified"))
		return
	}

	for _, option := range options {
        if option == "h" || option == "-help" {
            fmt.Print("Help message\n")
            return
        }
    }
	for _, option := range options {
		switch option {
		case "r":
			isRecursive = true
		case "f":
			isForced = true
		}
	}

	if len(names) > 1 {
		MultiRm(names)
	} else {
		SingleRm(names[0])
	}
}

func MultiRm(names []string) {
	for _, name := range names {
		SingleRm(name)
	}
}

func SingleRm(name string) {
	var absName = name
	if !filepath.IsAbs(name) {
		absName = filepath.Join(superpath.GetCurrentDirectory(), name)
	}

	fileInfo, err := os.Stat(absName)
	if err != nil {
		if os.IsNotExist(err) {
			errors.PrintError("rm", fmt.Errorf("%s: Error file or directory does not exist", name))
		}
		return
	}

	if fileInfo.IsDir() {
		if isForced {
			entries, err := os.ReadDir(absName)
			if err != nil {
				errors.PrintError("rm", fmt.Errorf("%s: Error checking directory content", name))
				return
			}
			isEmpty := len(entries) == 0

			if !isEmpty && !isForced{
				errors.PrintError("rm", fmt.Errorf("%s: Error directory is not empty", name))
				return
			}

			if isRecursive {
				err := os.RemoveAll(absName)
				if err != nil {
					errors.PrintError("rm", fmt.Errorf("%s: Error removing directory", name))
					return
				}
			} else {
				err := os.RemoveAll(absName)
				if err != nil {
					errors.PrintError("rm", fmt.Errorf("%s: Error removing directory", name))
					return
				}
			}
		} else {
			errors.PrintError("rm", fmt.Errorf("%s: Error cannot remove a directory", name))
			return
		}
		
	} else {
		err := os.Remove(absName)
		if err != nil {
			errors.PrintError("rm", fmt.Errorf("%s: Error removing file", name))
			return
		}
	}
}
