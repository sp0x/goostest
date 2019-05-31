package goostest

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print(runtime.GOOS)
}
