package errors

import (
	"fmt"
	"GoShell/colors"
)

func PrintError(cmd string, err error) {
	fmt.Printf(colors.Red + "GoShell: %s: %s\n" + colors.Reset, cmd, err)
}