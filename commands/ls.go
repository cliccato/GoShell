package commands

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"GoShell/colors"
	"GoShell/errors"
	"GoShell/superpath"
)

var ShowHidden = false

func Ls(directories []string, options []string) {
	for _, option := range options {
        if option == "h" || option == "-help" {
            fmt.Print("Help message\n")
            return
        }
    }
	for _, option := range options {
		switch option {
		case "a":
			ShowHidden = true
		}
	}

	if len(directories) > 1 {
		PrintMultiDir(directories)
	} else if len(directories) == 1 {
		PrintDir(directories[0])
	} else {
		PrintDir(".")
	}
}

func PrintMultiDir(directories []string) {
	for _, dir := range directories {
		var absDir string
		if !filepath.IsAbs(dir) {
			absDir = filepath.Join(superpath.GetCurrentDirectory(), dir)
		} else {
			absDir = dir
		}

		_, err := ioutil.ReadDir(absDir)
		if err != nil {
			errors.PrintError("ls", fmt.Errorf("%s: directory not found", dir))
			continue
		}

		fmt.Print(dir)
		PrintDir(dir)
		fmt.Print("\n")
	}
}

func PrintDir(dir string) {
	var absDir string
	if !filepath.IsAbs(dir) {
		absDir = filepath.Join(superpath.GetCurrentDirectory(), dir)
	} else {
		absDir = dir
	}

	files, err := ioutil.ReadDir(absDir)
	if err != nil {
		errors.PrintError("ls", fmt.Errorf("%s: directory not found", dir))
		return
	}

	for _, file := range files {
		if file.Name()[0] == '.' && !ShowHidden {
			continue
		} else if file.IsDir() {
			fmt.Println(colors.Blue + file.Name() + colors.Reset)
		} else if file.Mode().Perm()&0100 != 0 {
			fmt.Println(colors.Green + file.Name() + colors.Reset)
		} else {
			fmt.Println(file.Name())
		}
	}
}
