package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("offset exceeds file size", func(t *testing.T) {
		Copy("testdata/input.txt", "tmp/out.txt", 8000, 0)
		require.EqualError(t, ErrOffsetExceedsFileSize, "offset exceeds file size")
	})
}
