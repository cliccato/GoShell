package commands

import (
	"fmt"
)

func Clear() {
	fmt.Print("\033[2J\033[H")
}