package lib

import (
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
 * Converts given path to Posix (replacing \ with /) and removing ending slashes
 *
 * @param {string} givenPath Path to convert
 * @returns {string} Converted filepath
 */
func PosixifyPathNormalized(givenPath string) string {
	return strings.TrimRight(PosixifyPath(givenPath), "/")
}

