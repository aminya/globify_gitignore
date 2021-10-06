# globify-gitignore

Convert Gitignore to Glob patterns

A Go port of https://github.com/aminya/globify-gitignore

## Usage


```ts
import ("github.com/aminya/globify_gitignore")

GlobifyGitIgnoreFile(".") // path to a directory that has a .gitignore
```

You can use `globifyGitIgnore` to pass the gitignore content directly

```ts
import ("github.com/aminya/globify_gitignore")

func main() {
  gitignoreContent := `# OS metadata
  .DS_Store
  Thumbs.db

  # Node
  node_modules
  package-lock.json

  # TypeScript
  *.tsbuildinfo

  # Build directories
  dist
  `
  gitignoreDirectory = "./"

  globPatterns = globifyGitIgnore(gitignoreContent, gitignoreDirectory)
}
```

### API

These two functions are the main functions:

```go
/**
 * Parses and globifies the `.gitingore` file that exists in a directory
 *
 * @param {string} gitIgnoreDirectory The given directory that has the `.gitignore` file
 * @returns {([]string, error)} An array of glob patterns or an error if the file did not exist
 */
func GlobifyGitIgnoreFile(gitIgnoreDirectory string) ([]string, error)

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
) []string
```

### Other API

Other possibly useful functions:

```go
/**
 * Globify a path
 * @param {string} givenPath The given path to be globified
 * @param {Optional string} givenDirectory [process.cwd()] The cwd to use to resolve relative path names
 * @returns {Promise<string | [string, string]>} The glob path or the file path itself
 */
func GlobifyPath(
	givenPath string,
	givenDirectory ...string,
) []string


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
) []string

/**
 * Globifies a directory
 *
 * @param {string} givenDirectory The given directory to be globified
 */
func GlobifyDirectory(givenDirectory string) string

/// Is this string a valid path
func IsPath(path string, extended bool) bool

/// Is this string an invalid path?
func IsInvalidPath(path string, extended bool) bool

/**
 * Converts given path to Posix (replacing \ with /)
 *
 * @param {string} givenPath Path to convert
 * @returns {string} Converted filepath
 */
func PosixifyPath(givenPath string) string

/**
 * Converts given path to Posix (replacing \ with /)
 *
 * @param {string} givenPath Path to convert
 * @returns {string} Converted filepath
 */
func PosixifyPath(givenPath string) string

/**
 * Removes the ending slash from the given path
 *
 * @param {string} givenPath Path to convert
 * @returns {string} Converted filepath
 */
func RemoveEndingSlash(givenPath string) string

/**
 * A line starting with # serves as a comment. Put a backslash ("") in front of the first hash for patterns that begin
 * with a hash.
 */
func IsGitIgnoreComment(pattern string) bool

/** Trailing spaces should be removed unless they are quoted with backslash ("\ "). */
func TrimTrailingWhitespace(str string) string


/** Remove leading whitespace */
func TrimLeadingWhiteSpace(str string) string

/** Remove whitespace from a gitignore entry */
func TrimWhiteSpace(str string) string

/**
 * Get the type of the given path
 *
 * @param {string} givenPath Absolute path
 * @returns {PathType}
 */
func GetPathType(filepath string) PathType
```

## Contributing

- Let me know if you encounter any bugs.
- Feature requests are always welcome.