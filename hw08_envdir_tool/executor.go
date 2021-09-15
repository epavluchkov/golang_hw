package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for key, val := range env {
		os.Unsetenv(key)
		if !val.NeedRemove {
			os.Setenv(key, val.Value)
		}
	}

	com := exec.Command(cmd[0], cmd[1:]...) // #nosec G204
	com.Stdin = os.Stdin
	com.Stdout = os.Stdout
	com.Stderr = os.Stderr

	if err := com.Run(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
