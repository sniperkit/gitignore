package dart

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# See https://www.dartlang.org/tools/private-files.html",
		"",
		"# Files and directories created by pub",
		".buildlog",
		".packages",
		".project",
		".pub/",
		"build/",
		"**/packages/",
		"",
		"# Files created by dart2js",
		"# (Most Dart developers will use pub build to compile Dart, use/modify these",
		"#  rules if you intend to use dart2js directly",
		"#  Convention is to use extension '.dart.js' for Dart compiled to Javascript to",
		"#  differentiate from explicit Javascript files)",
		"*.dart.js",
		"*.part.js",
		"*.js.deps",
		"*.js.map",
		"*.info.json",
		"",
		"# Directory created by dartdoc",
		"doc/api/",
		"",
		"# Don't commit pubspec lock file",
		"# (Library packages only! Remove pattern if developing an application package)",
		"pubspec.lock",
		"",
	}
	template.Add("github/dart", strings.Join(ignore, "\n"))
}
