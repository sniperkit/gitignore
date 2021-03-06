package rust

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Generated by Cargo",
		"# will have compiled files and executables",
		"/target/",
		"",
		"# Remove Cargo.lock from gitignore if creating an executable, leave it for libraries",
		"# More information here http://doc.crates.io/guide.html#cargotoml-vs-cargolock",
		"Cargo.lock",
		"",
	}
	template.Add("github/rust", strings.Join(ignore, "\n"))
}
