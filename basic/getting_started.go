package basic

import (
	"fmt"
	"runtime"
)

/*
1. printing the go version by go program
*/
// PrintGoEnv prints the go environment variables
func PrintGoEnv() {
	fmt.Printf("Go version : %s\n", runtime.Version())
	fmt.Printf("Go OS : %s\n",runtime.GOOS)
	fmt.Printf("Go Architecture : %s\n",runtime.GOARCH)
	fmt.Printf("Go Root : %s\n",runtime.GOROOT())
}
