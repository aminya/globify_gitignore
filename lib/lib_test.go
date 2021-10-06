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

	assert.Equal(t, isGitIgnoreComment("# aaa"), true)
	assert.Equal(t, isGitIgnoreComment("#aa"), true)
	assert.Equal(t, isGitIgnoreComment(" #"), false)
	assert.Equal(t, isGitIgnoreComment(" "), false)
	assert.Equal(t, isGitIgnoreComment("aa"), false)

	assert.Equal(t, trimTrailingWhitespace("aa  "), "aa")
	assert.Equal(t, trimTrailingWhitespace("aa \\ "), "aa  ")
	assert.Equal(t, trimTrailingWhitespace("aa \\  "), "aa   ")
	assert.Equal(t, trimTrailingWhitespace("aa"), "aa")

	assert.Equal(t, trimLeadingWhiteSpace("aa"), "aa")
	assert.Equal(t, trimLeadingWhiteSpace("  aa"), "aa")
	assert.Equal(t, trimLeadingWhiteSpace(" \\ aa"), "\\ aa")

	assert.Equal(t, trimWhiteSpace("aa  "), "aa")
	assert.Equal(t, trimWhiteSpace("aa \\ "), "aa  ")
	assert.Equal(t, trimWhiteSpace("aa \\  "), "aa   ")
	assert.Equal(t, trimWhiteSpace("aa"), "aa")
	assert.Equal(t, trimWhiteSpace("  aa"), "aa")
	assert.Equal(t, trimWhiteSpace(" \\ aa"), "\\ aa")
	assert.Equal(t, trimWhiteSpace("  aa  "), "aa")
	assert.Equal(t, trimWhiteSpace("  aa \\ "), "aa  ")
	assert.Equal(t, trimWhiteSpace("  aa \\  "), "aa   ")
	assert.Equal(t, trimWhiteSpace("  aa"), "aa")
	assert.Equal(t, trimWhiteSpace("  \\ aa"), "\\ aa")
}

func TestGlobifyGitIgnoreEntry(t *testing.T) {
	// Files or directories
	assert.Equal(t, GlobifyGitIgnoreEntry("dir_or_file"), []string{"!**/dir_or_file", "!**/dir_or_file/**"})

	// Relative dir
	assert.Equal(t, GlobifyGitIgnoreEntry("dir/"), []string{"!dir/**"})

	// Absolute paths
	assert.Equal(t, GlobifyGitIgnoreEntry("/abs_dir_or_file"), []string{"!abs_dir_or_file", "!abs_dir_or_file/**"})
	assert.Equal(t, GlobifyGitIgnoreEntry("/abs_dir/abs_dir_or_file"), []string{"!abs_dir/abs_dir_or_file", "!abs_dir/abs_dir_or_file/**"})
	assert.Equal(t, GlobifyGitIgnoreEntry("/abs_dir/abs_dir/"), []string{"!abs_dir/abs_dir/", "!abs_dir/abs_dir//**"})
	assert.Equal(t, GlobifyGitIgnoreEntry("C:/abs_dir_or_file"), []string{"!C:/abs_dir_or_file", "!C:/abs_dir_or_file/**"})
	assert.Equal(t, GlobifyGitIgnoreEntry("C:/abs_dir/abs_dir_or_file"), []string{"!C:/abs_dir/abs_dir_or_file", "!C:/abs_dir/abs_dir_or_file/**"})
	assert.Equal(t, GlobifyGitIgnoreEntry("C:/abs_dir/abs_dir/"), []string{"!C:/abs_dir/abs_dir/", "!C:/abs_dir/abs_dir//**"})
}
