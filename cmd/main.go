package main

import (
	"fmt"
	"github.com/sp0x/goostest/tools"
	"runtime"
)

func main() {
	fmt.Printf("Os is: %v", runtime.GOOS)
	if tools.HasBash() {
		fmt.Println("Bash is found!")
	} else {
		fmt.Println("Couldn't find bash")
	}
}
