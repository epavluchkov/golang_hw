package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("read files", func(t *testing.T) {
		pwd, _ := os.Getwd()
		env, err := ReadDir(pwd + "/testdata/env")
		require.Nil(t, err)
		require.Equal(t, "bar", env["BAR"].Value)
		require.Equal(t, true, env["UNSET"].NeedRemove)
	})
}
