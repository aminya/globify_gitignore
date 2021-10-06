package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosixifyPath(t *testing.T) {
	assert.Equal(t, PosixifyPath("C:\\hey"), "C:/hey")
	assert.Equal(t, PosixifyPath("hey/"), "hey/")
}

func TestRemoveEndingSlash(t *testing.T) {
	assert.Equal(t, RemoveEndingSlash(PosixifyPath("C:\\hey\\")), "C:/hey")
	assert.Equal(t, RemoveEndingSlash(PosixifyPath("hey/")), "hey")
}

func TestIsEmptyLine(t *testing.T) {
	assert.Equal(t, IsEmptyLine("  "), true)
	assert.Equal(t, IsEmptyLine(" "), true)
	assert.Equal(t, IsEmptyLine(" \n"), true)
	assert.Equal(t, IsEmptyLine(" #"), false)
}

func TestTrimTrailingWhitespace(t *testing.T) {
	assert.Equal(t, TrimTrailingWhitespace("aa  "), "aa")
	assert.Equal(t, TrimTrailingWhitespace("aa \\ "), "aa  ")
	assert.Equal(t, TrimTrailingWhitespace("aa \\  "), "aa   ")
	assert.Equal(t, TrimTrailingWhitespace("aa"), "aa")
}

func TestTrimLeadingWhiteSpace(t *testing.T) {
	assert.Equal(t, TrimLeadingWhiteSpace("aa"), "aa")
	assert.Equal(t, TrimLeadingWhiteSpace("  aa"), "aa")
	assert.Equal(t, TrimLeadingWhiteSpace(" \\ aa"), "\\ aa")
}

func TestTrimWhiteSpace(t *testing.T) {
	assert.Equal(t, TrimWhiteSpace("aa  "), "aa")
	assert.Equal(t, TrimWhiteSpace("aa \\ "), "aa  ")
	assert.Equal(t, TrimWhiteSpace("aa \\  "), "aa   ")
	assert.Equal(t, TrimWhiteSpace("aa"), "aa")
	assert.Equal(t, TrimWhiteSpace("  aa"), "aa")
	assert.Equal(t, TrimWhiteSpace(" \\ aa"), "\\ aa")
	assert.Equal(t, TrimWhiteSpace("  aa  "), "aa")
	assert.Equal(t, TrimWhiteSpace("  aa \\ "), "aa  ")
	assert.Equal(t, TrimWhiteSpace("  aa \\  "), "aa   ")
	assert.Equal(t, TrimWhiteSpace("  aa"), "aa")
	assert.Equal(t, TrimWhiteSpace("  \\ aa"), "\\ aa")
}

func TestGlobifyDirectory(t *testing.T) {
	assert.Equal(t, GlobifyDirectory("hey/"), "hey/**")
	assert.Equal(t, GlobifyDirectory("/home/"), "/home/**")
	assert.Equal(t, GlobifyDirectory("C:\\hey\\"), "C:/hey/**")
	assert.Equal(t, GlobifyDirectory("hey\\hey2\\"), "hey/hey2/**")
}

func TestIsGitIgnoreComment(t *testing.T) {
	assert.Equal(t, IsGitIgnoreComment("# aaa"), true)
	assert.Equal(t, IsGitIgnoreComment("#aa"), true)
	assert.Equal(t, IsGitIgnoreComment(" #"), false)
	assert.Equal(t, IsGitIgnoreComment(" "), false)
	assert.Equal(t, IsGitIgnoreComment("aa"), false)
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
