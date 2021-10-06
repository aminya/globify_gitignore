package lib

import (
	"regexp"
	"strings"
)

/**
 * Converts given path to Posix (replacing \ with /)
 *
 * @param {string} givenPath Path to convert
 * @returns {string} Converted filepath
 */
func PosixifyPath(givenPath string) string {
	return strings.ReplaceAll(givenPath, "\\", "/")
}

/**
 * Removes the ending slash from the given path
 *
 * @param {string} givenPath Path to convert
 * @returns {string} Converted filepath
 */
func RemoveEndingSlash(givenPath string) string {
	return strings.TrimRight(givenPath, "/")
}

/**
 * Globifies a directory
 *
 * @param {string} givenDirectory The given directory to be globified
 */
func GlobifyDirectory(givenDirectory string) string {
	return RemoveEndingSlash(PosixifyPath(givenDirectory)) + "/**"
}

func isEmptyLine(str string) bool {
	whiteSpaceRegex := regexp.MustCompile(`^\s*$`)
	return whiteSpaceRegex.MatchString(str)
}

/**
 * A line starting with # serves as a comment. Put a backslash ("") in front of the first hash for patterns that begin
 * with a hash.
 */
func isGitIgnoreComment(pattern string) bool {
	return pattern[0] == '#'
}

/** Trailing spaces should be removed unless they are quoted with backslash ("\ "). */
func trimTrailingWhitespace(str string) string {
	escaped_trailing_whitespace := regexp.MustCompile(`\\\s+$`)
	if !escaped_trailing_whitespace.MatchString(str) {
		trailing_whitespace := regexp.MustCompile(`\s+$`)
		// No escaped trailing whitespace, remove
		return trailing_whitespace.ReplaceAllString(str, "")
	} else {
		// Trailing whitespace detected, remove only the backslash
		backslash := regexp.MustCompile(`\\(\s+)$`)
		return backslash.ReplaceAllString(str, "$1")
	}
}

/** Remove leading whitespace */
func trimLeadingWhiteSpace(str string) string {
	leading_whitespace := regexp.MustCompile(`^\s+`)
	return leading_whitespace.ReplaceAllString(str, "")
}

/** Remove whitespace from a gitignore entry */
func trimWhiteSpace(str string) string {
	return trimLeadingWhiteSpace(trimTrailingWhitespace(str))
}

/** Enum that specifies the path type. 0 for file, 1 for directory, 2 for others */
type PathType uint

const (
	PathTypeFile      PathType = 0
	PathTypeDirectory PathType = 1
	PathTypeOther     PathType = 2
)

/**
 * Get the type of the given path
 *
 * @param {string} givenPath Absolute path
 * @returns {PathType}
 */
func getPathType(filepath string) PathType {
	pathStat, err := os.Lstat(filepath)
	if err != nil {
		return PathTypeOther
	}
	switch mode := pathStat.Mode(); {
	case mode.IsRegular():
		return PathTypeFile
	case mode.IsDir():
		return PathTypeDirectory
	case mode&fs.ModeSymlink != 0:
		return PathTypeOther
	case mode&fs.ModeNamedPipe != 0:
		return PathTypeOther
	default:
		return PathTypeOther
	}
}
