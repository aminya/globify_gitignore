package lib

import (
	"log"
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

func TestGlobifyGitIgnore(t *testing.T) {

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

*.dll
*.exe
*.cmd
*.pdb
*.suo
*.js
*.user
*.cache
*.cs
*.sln
*.csproj
*.map
*.swp
*.code-workspace
*.log
.DS_Store

_Resharper.DefinitelyTyped
bin
obj
Properties

# VIM backup files
*~

# test folder
_infrastructure/tests/build

# IntelliJ based IDEs
.idea
*.iml

*.js.map
!*.js/
!scripts/new-package.js
!scripts/not-needed.js
!scripts/lint.js

# npm
node_modules
package-lock.json
npm-debug.log

# Sublime
.sublimets

# Visual Studio Code
.settings/launch.json
.vs
.vscode
.history

# yarn
yarn.lock

# pnpm
shrinkwrap.yaml
pnpm-lock.yaml
pnpm-debug.log

# Output of 'npm pack'
*.tgz
`
	assert.Equal(t, GlobifyGitIgnore(gitignoreContent, "./fixtures"), []string{
		`!./fixtures/**/.DS_Store`,
		`!./fixtures/**/Thumbs.db`,
		`!./fixtures/**/node_modules`,
		`!./fixtures/**/package-lock.json`,
		`!./fixtures/**/*.tsbuildinfo`,
		`!./fixtures/**/dist`,
		`!./fixtures/**/*.dll`,
		`!./fixtures/**/*.exe`,
		`!./fixtures/**/*.cmd`,
		`!./fixtures/**/*.pdb`,
		`!./fixtures/**/*.suo`,
		`!./fixtures/**/*.js`,
		`!./fixtures/**/*.user`,
		`!./fixtures/**/*.cache`,
		`!./fixtures/**/*.cs`,
		`!./fixtures/**/*.sln`,
		`!./fixtures/**/*.csproj`,
		`!./fixtures/**/*.map`,
		`!./fixtures/**/*.swp`,
		`!./fixtures/**/*.code-workspace`,
		`!./fixtures/**/*.log`,
		`!./fixtures/**/_Resharper.DefinitelyTyped`,
		`!./fixtures/**/bin`,
		`!./fixtures/**/obj`,
		`!./fixtures/**/Properties`,
		`!./fixtures/**/*~`,
		`!./fixtures/_infrastructure/tests/build`,
		`!./fixtures/**/.idea`,
		`!./fixtures/**/*.iml`,
		`!./fixtures/**/*.js.map`,
		`./fixtures/*.js/**`,
		`./fixtures/scripts/new-package.js`,
		`./fixtures/scripts/not-needed.js`,
		`./fixtures/scripts/lint.js`,
		`!./fixtures/**/npm-debug.log`,
		`!./fixtures/**/.sublimets`,
		`!./fixtures/.settings/launch.json`,
		`!./fixtures/**/.vs`,
		`!./fixtures/**/.vscode`,
		`!./fixtures/**/.history`,
		`!./fixtures/**/yarn.lock`,
		`!./fixtures/**/shrinkwrap.yaml`,
		`!./fixtures/**/pnpm-lock.yaml`,
		`!./fixtures/**/pnpm-debug.log`,
		`!./fixtures/**/*.tgz`,
		`!./fixtures/**/.DS_Store/**`,
		`!./fixtures/**/Thumbs.db/**`,
		`!./fixtures/**/node_modules/**`,
		`!./fixtures/**/package-lock.json/**`,
		`!./fixtures/**/*.tsbuildinfo/**`,
		`!./fixtures/**/dist/**`,
		`!./fixtures/**/*.dll/**`,
		`!./fixtures/**/*.exe/**`,
		`!./fixtures/**/*.cmd/**`,
		`!./fixtures/**/*.pdb/**`,
		`!./fixtures/**/*.suo/**`,
		`!./fixtures/**/*.js/**`,
		`!./fixtures/**/*.user/**`,
		`!./fixtures/**/*.cache/**`,
		`!./fixtures/**/*.cs/**`,
		`!./fixtures/**/*.sln/**`,
		`!./fixtures/**/*.csproj/**`,
		`!./fixtures/**/*.map/**`,
		`!./fixtures/**/*.swp/**`,
		`!./fixtures/**/*.code-workspace/**`,
		`!./fixtures/**/*.log/**`,
		`!./fixtures/**/_Resharper.DefinitelyTyped/**`,
		`!./fixtures/**/bin/**`,
		`!./fixtures/**/obj/**`,
		`!./fixtures/**/Properties/**`,
		`!./fixtures/**/*~/**`,
		`!./fixtures/_infrastructure/tests/build/**`,
		`!./fixtures/**/.idea/**`,
		`!./fixtures/**/*.iml/**`,
		`!./fixtures/**/*.js.map/**`,
		`./fixtures/scripts/new-package.js/**`,
		`./fixtures/scripts/not-needed.js/**`,
		`./fixtures/scripts/lint.js/**`,
		`!./fixtures/**/npm-debug.log/**`,
		`!./fixtures/**/.sublimets/**`,
		`!./fixtures/.settings/launch.json/**`,
		`!./fixtures/**/.vs/**`,
		`!./fixtures/**/.vscode/**`,
		`!./fixtures/**/.history/**`,
		`!./fixtures/**/yarn.lock/**`,
		`!./fixtures/**/shrinkwrap.yaml/**`,
		`!./fixtures/**/pnpm-lock.yaml/**`,
		`!./fixtures/**/pnpm-debug.log/**`,
		`!./fixtures/**/*.tgz/**`,
	})

	assert.Equal(t, GlobifyGitIgnore(gitignoreContent), []string{
		`!**/.DS_Store`,
		`!**/Thumbs.db`,
		`!**/node_modules`,
		`!**/package-lock.json`,
		`!**/*.tsbuildinfo`,
		`!**/dist`,
		`!**/*.dll`,
		`!**/*.exe`,
		`!**/*.cmd`,
		`!**/*.pdb`,
		`!**/*.suo`,
		`!**/*.js`,
		`!**/*.user`,
		`!**/*.cache`,
		`!**/*.cs`,
		`!**/*.sln`,
		`!**/*.csproj`,
		`!**/*.map`,
		`!**/*.swp`,
		`!**/*.code-workspace`,
		`!**/*.log`,
		`!**/_Resharper.DefinitelyTyped`,
		`!**/bin`,
		`!**/obj`,
		`!**/Properties`,
		`!**/*~`,
		`!_infrastructure/tests/build`,
		`!**/.idea`,
		`!**/*.iml`,
		`!**/*.js.map`,
		`..js/**`,
		`.cripts/new-package.js`,
		`.cripts/not-needed.js`,
		`.cripts/lint.js`,
		`!**/npm-debug.log`,
		`!**/.sublimets`,
		`!.settings/launch.json`,
		`!**/.vs`,
		`!**/.vscode`,
		`!**/.history`,
		`!**/yarn.lock`,
		`!**/shrinkwrap.yaml`,
		`!**/pnpm-lock.yaml`,
		`!**/pnpm-debug.log`,
		`!**/*.tgz`,
		`!**/.DS_Store/**`,
		`!**/Thumbs.db/**`,
		`!**/node_modules/**`,
		`!**/package-lock.json/**`,
		`!**/*.tsbuildinfo/**`,
		`!**/dist/**`,
		`!**/*.dll/**`,
		`!**/*.exe/**`,
		`!**/*.cmd/**`,
		`!**/*.pdb/**`,
		`!**/*.suo/**`,
		`!**/*.js/**`,
		`!**/*.user/**`,
		`!**/*.cache/**`,
		`!**/*.cs/**`,
		`!**/*.sln/**`,
		`!**/*.csproj/**`,
		`!**/*.map/**`,
		`!**/*.swp/**`,
		`!**/*.code-workspace/**`,
		`!**/*.log/**`,
		`!**/_Resharper.DefinitelyTyped/**`,
		`!**/bin/**`,
		`!**/obj/**`,
		`!**/Properties/**`,
		`!**/*~/**`,
		`!_infrastructure/tests/build/**`,
		`!**/.idea/**`,
		`!**/*.iml/**`,
		`!**/*.js.map/**`,
		`.cripts/new-package.js/**`,
		`.cripts/not-needed.js/**`,
		`.cripts/lint.js/**`,
		`!**/npm-debug.log/**`,
		`!**/.sublimets/**`,
		`!.settings/launch.json/**`,
		`!**/.vs/**`,
		`!**/.vscode/**`,
		`!**/.history/**`,
		`!**/yarn.lock/**`,
		`!**/shrinkwrap.yaml/**`,
		`!**/pnpm-lock.yaml/**`,
		`!**/pnpm-debug.log/**`,
		`!**/*.tgz/**`,
	})

}

func TestGlobifyGitIgnoreFile(t *testing.T) {
	globs, err := GlobifyGitIgnoreFile("./fixtures")
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, globs, []string{
		`!./fixtures/**/.DS_Store`,
		`!./fixtures/**/Thumbs.db`,
		`!./fixtures/**/node_modules`,
		`!./fixtures/**/package-lock.json`,
		`!./fixtures/**/*.tsbuildinfo`,
		`!./fixtures/**/dist`,
		`!./fixtures/**/*.dll`,
		`!./fixtures/**/*.exe`,
		`!./fixtures/**/*.cmd`,
		`!./fixtures/**/*.pdb`,
		`!./fixtures/**/*.suo`,
		`!./fixtures/**/*.js`,
		`!./fixtures/**/*.user`,
		`!./fixtures/**/*.cache`,
		`!./fixtures/**/*.cs`,
		`!./fixtures/**/*.sln`,
		`!./fixtures/**/*.csproj`,
		`!./fixtures/**/*.map`,
		`!./fixtures/**/*.swp`,
		`!./fixtures/**/*.code-workspace`,
		`!./fixtures/**/*.log`,
		`!./fixtures/**/_Resharper.DefinitelyTyped`,
		`!./fixtures/**/bin`,
		`!./fixtures/**/obj`,
		`!./fixtures/**/Properties`,
		`!./fixtures/**/*~`,
		`!./fixtures/_infrastructure/tests/build`,
		`!./fixtures/**/.idea`,
		`!./fixtures/**/*.iml`,
		`!./fixtures/**/*.js.map`,
		`./fixtures/*.js/**`,
		`./fixtures/scripts/new-package.js`,
		`./fixtures/scripts/not-needed.js`,
		`./fixtures/scripts/lint.js`,
		`!./fixtures/**/npm-debug.log`,
		`!./fixtures/**/.sublimets`,
		`!./fixtures/.settings/launch.json`,
		`!./fixtures/**/.vs`,
		`!./fixtures/**/.vscode`,
		`!./fixtures/**/.history`,
		`!./fixtures/**/yarn.lock`,
		`!./fixtures/**/shrinkwrap.yaml`,
		`!./fixtures/**/pnpm-lock.yaml`,
		`!./fixtures/**/pnpm-debug.log`,
		`!./fixtures/**/*.tgz`,
		`!./fixtures/**/.DS_Store/**`,
		`!./fixtures/**/Thumbs.db/**`,
		`!./fixtures/**/node_modules/**`,
		`!./fixtures/**/package-lock.json/**`,
		`!./fixtures/**/*.tsbuildinfo/**`,
		`!./fixtures/**/dist/**`,
		`!./fixtures/**/*.dll/**`,
		`!./fixtures/**/*.exe/**`,
		`!./fixtures/**/*.cmd/**`,
		`!./fixtures/**/*.pdb/**`,
		`!./fixtures/**/*.suo/**`,
		`!./fixtures/**/*.js/**`,
		`!./fixtures/**/*.user/**`,
		`!./fixtures/**/*.cache/**`,
		`!./fixtures/**/*.cs/**`,
		`!./fixtures/**/*.sln/**`,
		`!./fixtures/**/*.csproj/**`,
		`!./fixtures/**/*.map/**`,
		`!./fixtures/**/*.swp/**`,
		`!./fixtures/**/*.code-workspace/**`,
		`!./fixtures/**/*.log/**`,
		`!./fixtures/**/_Resharper.DefinitelyTyped/**`,
		`!./fixtures/**/bin/**`,
		`!./fixtures/**/obj/**`,
		`!./fixtures/**/Properties/**`,
		`!./fixtures/**/*~/**`,
		`!./fixtures/_infrastructure/tests/build/**`,
		`!./fixtures/**/.idea/**`,
		`!./fixtures/**/*.iml/**`,
		`!./fixtures/**/*.js.map/**`,
		`./fixtures/scripts/new-package.js/**`,
		`./fixtures/scripts/not-needed.js/**`,
		`./fixtures/scripts/lint.js/**`,
		`!./fixtures/**/npm-debug.log/**`,
		`!./fixtures/**/.sublimets/**`,
		`!./fixtures/.settings/launch.json/**`,
		`!./fixtures/**/.vs/**`,
		`!./fixtures/**/.vscode/**`,
		`!./fixtures/**/.history/**`,
		`!./fixtures/**/yarn.lock/**`,
		`!./fixtures/**/shrinkwrap.yaml/**`,
		`!./fixtures/**/pnpm-lock.yaml/**`,
		`!./fixtures/**/pnpm-debug.log/**`,
		`!./fixtures/**/*.tgz/**`,
	})
}
