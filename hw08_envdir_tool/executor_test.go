package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("simple test with nil environment", func(t *testing.T) {
		pwd, _ := os.Getwd()
		res := RunCmd([]string{"/bin/bash", pwd + "/testdata/echo.sh"}, nil)
		require.Equal(t, 0, res)
	})
}
