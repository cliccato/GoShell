package commands

import (
	"GoShell/superpath"
	"fmt"
)

func Pwd() {
	fmt.Println(superpath.GetCurrentDirectory())
}
