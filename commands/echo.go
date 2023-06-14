package commands

import (
	"fmt"
	"strings"
	"GoShell/errors"
)

func Echo(params []string) {
	if len(params) > 0 {
		fmt.Println(strings.Join(params, " "))
	} else {
		errors.PrintError("Echo", fmt.Errorf("No params specified"))
	}
}
