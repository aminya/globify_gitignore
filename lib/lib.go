package lib

import (
	"io/fs"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/lithammer/dedent"
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

func IsEmptyLine(str string) bool {
	whiteSpaceRegex := regexp.MustCompile(`^\s*$`)
	return whiteSpaceRegex.MatchString(str)
}

/**
 * A line starting with # serves as a comment. Put a backslash ("") in front of the first hash for patterns that begin
 * with a hash.
 */
func IsGitIgnoreComment(pattern string) bool {
	return pattern[0] == '#'
}

/** Trailing spaces should be removed unless they are quoted with backslash ("\ "). */
func TrimTrailingWhitespace(str string) string {
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
func TrimLeadingWhiteSpace(str string) string {
	leading_whitespace := regexp.MustCompile(`^\s+`)
	return leading_whitespace.ReplaceAllString(str, "")
}

/** Remove whitespace from a gitignore entry */
func TrimWhiteSpace(str string) string {
	return TrimLeadingWhiteSpace(TrimTrailingWhitespace(str))
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
func GetPathType(filepath string) PathType {
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

func isInvalidPath(path string, extended bool) bool {
	/*
	 * Go port of
	 * is-invalid-path <https://github.com/jonschlinkert/is-invalid-path>
	 *
	 * Copyright (c) 2015-2018, Jon Schlinkert.
	 * Released under the MIT License.
	 */

	if path == "" {
		return true
	}

	// https://msdn.microsoft.com/en-us/library/windows/desktop/aa365247(v=vs.85).aspx#maxpath
	MAX_PATH := 260
	if extended {
		MAX_PATH = 32767
	}

	if len(path) > (MAX_PATH - 12) {
		return true
	}

	// TODO
	// const rootPath = path.parse(path).root
	// if rootPath {
	// 	path = path.slice(rootPath.length)
	// }

	// https://msdn.microsoft.com/en-us/library/windows/desktop/aa365247(v=vs.85).aspx#Naming_Conventions
	invalidFileRegex := regexp.MustCompile(`[<>:"|?*]`)
	return invalidFileRegex.MatchString(path)
}

func isPath(path string, extended bool) bool {
	return !isInvalidPath(path, extended)
}

/**
 * @param {string} gitIgnoreEntry One git ignore entry
 * @param {Optional string} gitIgnoreDirectory The directory of gitignore
 * @returns {[string] | [string, string]} The equivalent glob
 *
 * NOTE: it expects a **valid** non-comment git-ignore entry  with no surrounding whitespace.
 * NOTE: Gitignore expects that paths are posixified. So, if you are passing Windows path to this function directly without poxifying them (using {PosixifyPath}), you are passing invalid gitignore entry, and so you will get invalid Glob pattern.
 */
func GlobifyGitIgnoreEntry(
	gitIgnoreEntry string,
	gitIgnoreDirectory ...string,
) []string {
	// output glob entry
	entry := gitIgnoreEntry
	// Process the entry beginning
	// '!' in .gitignore means to force include the pattern
	// remove "!" to allow the processing of the pattern and swap ! in the end of the loop
	forceInclude := false

	hasGitIgnoreDirectory := len(gitIgnoreDirectory) == 1 // TODO find a better way for optional arguments in Go

	if entry[0] == '!' {
		entry = entry[1:]
		forceInclude = true
	}

	// If there is a separator at the beginning or middle (or both) of the pattern,
	// then the pattern is relative to the directory level of the particular .gitignore file itself
	// Process slash

	pathType := PathTypeOther

	if entry[0] == '/' {
		// Patterns starting with '/' in gitignore are considered relative to the project directory while glob
		// treats them as relative to the OS root directory.
		// So we trim the slash to make it relative to project folder from glob perspective.
		entry = entry[1:]

		// Check if it is a directory or file
		if isPath(entry, true) {
			if hasGitIgnoreDirectory {
				pathType = GetPathType(path.Join(gitIgnoreDirectory[0], entry))
			} else {
				pathType = GetPathType(entry)
			}
		}
	} else {
		slashPlacement := strings.Index(entry, "/")

		if slashPlacement == -1 {
			// Patterns that don't have `/` are '**/' from glob perspective (can match at any level)
			if !strings.HasPrefix(entry, "**/") {
				entry = "**/" + entry
			}
		} else if slashPlacement == len(entry)-1 {
			// If there is a separator at the end of the pattern then it only matches directories
			// slash is in the end
			pathType = PathTypeDirectory
		} else {
			// has `/` in the middle so it is a relative path
			// Check if it is a directory or file
			if isPath(entry, true) {
				if hasGitIgnoreDirectory {
					pathType = GetPathType(path.Join(gitIgnoreDirectory[0], entry))
				} else {
					pathType = GetPathType(entry)
				}
			}
		}
	}

	// prepend the absolute root directory
	if hasGitIgnoreDirectory {
		entry = PosixifyPath(gitIgnoreDirectory[0]) + "/" + entry
	}

	// swap !
	if !(forceInclude) {
		entry = "!" + entry
	}

	// TODO use a tagged union instead of an array?

	// Process the entry ending
	if pathType == PathTypeDirectory {
		// in glob this is equal to `directory/**`
		if strings.HasSuffix(entry, "/") {
			return []string{entry + "**"}
		} else {
			return []string{entry + "/**"}
		}
	} else if pathType == PathTypeFile {
		// return as is for file
		return []string{entry}
	} else if !strings.HasSuffix(entry, "/**") {
		// the pattern can match both files and directories
		// so we should include both `entry` and `entry/**`
		content := entry + "/**"
		return []string{entry, content}
	} else {
		return []string{entry}
	}
}

/**
 * Globify the content of a `.gitignore` file
 *
 * @param {string} gitIgnoreContent The content of the gitignore file
 * @param {Optional string} gitIgnoreDirectory The directory of gitignore
 * @returns {[]string} An array of glob patterns
 */
func GlobifyGitIgnore(
	gitIgnoreContent string,
	gitIgnoreDirectory ...string,
) []string {
	gitIgnoreContentDedented := dedent.Dedent(gitIgnoreContent)
	gitIgnoreContentLines := strings.Split(gitIgnoreContentDedented, "\n")

	gitIgnoreEntries := []string{}
	for iLine := range gitIgnoreContentLines {
		entry := gitIgnoreContentLines[iLine]
		// Exclude empty lines and comments (filtering).
		if !(IsEmptyLine(entry) || IsGitIgnoreComment(entry)) {
			// Remove surrounding whitespace
			entryTrimmed := TrimWhiteSpace(entry)

			// out
			gitIgnoreEntries = append(gitIgnoreEntries, entryTrimmed)
		}
	}
	gitIgnoreEntriesNum := len(gitIgnoreEntries)

	globEntries := []string{} // TODO reserve at least gitIgnoreEntriesNum?

	for iEntry := 0; iEntry < gitIgnoreEntriesNum; iEntry++ {

		globifyOutput := GlobifyGitIgnoreEntry(gitIgnoreEntries[iEntry], gitIgnoreDirectory...)

		// Check if `GlobifyGitIgnoreEntry` returns a pair or a string
		if len(globifyOutput) == 1 {
			// string
			globEntries = append(globEntries, globifyOutput[0]) // Place the entry in the output array
		} else {
			// pair
			globEntries = append(globEntries,
				globifyOutput[0], // Place the entry in the output array
				globifyOutput[1]) // Push the additional entry
		}
	}

	// TODO unique in the end?
	return globEntries
}
