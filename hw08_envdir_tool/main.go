package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	env, err := ReadDir(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if res := RunCmd(args[2:], env); res != 0 {
		fmt.Println("some error")
	}
}
