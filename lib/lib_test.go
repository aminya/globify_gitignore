package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	assert.Equal(t, PosixifyPath("C:\\hey"), "C:/hey")
	assert.Equal(t, PosixifyPath("hey/"), "hey/")

	assert.Equal(t, RemoveEndingSlash(PosixifyPath("C:\\hey\\")), "C:/hey")
	assert.Equal(t, RemoveEndingSlash(PosixifyPath("hey/")), "hey")

	assert.Equal(t, GlobifyDirectory("hey/"), "hey/**")
	assert.Equal(t, GlobifyDirectory("/home/"), "/home/**")
	assert.Equal(t, GlobifyDirectory("C:\\hey\\"), "C:/hey/**")
	assert.Equal(t, GlobifyDirectory("hey\\hey2\\"), "hey/hey2/**")

	assert.Equal(t, isEmptyLine("  "), true)
	assert.Equal(t, isEmptyLine(" "), true)
	assert.Equal(t, isEmptyLine(" \n"), true)
	assert.Equal(t, isEmptyLine(" #"), false)

}
