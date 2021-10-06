package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosixifyPath(t *testing.T) {
	assert.Equal(t, PosixifyPath("C:\\hey"), "C:/hey")
	assert.Equal(t, PosixifyPath("hey/"), "hey/")
	assert.Equal(t, RemoveEndingSlash(PosixifyPath("C:\\hey\\")), "C:/hey")
	assert.Equal(t, RemoveEndingSlash(PosixifyPath("hey/")), "hey")
}
