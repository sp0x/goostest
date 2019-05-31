package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	if tools.HasBash() {
		fmt.Println("Bash is found!")
	} else {
		fmt.Println("Couldn't find bash")
	}
}
