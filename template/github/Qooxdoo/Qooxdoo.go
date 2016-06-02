package Qooxdoo

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"cache",
		"cache-downloads",
		"inspector",
		"api",
		"source/inspector.html",
	}
	template.SetTemplate("GitHub/Qooxdoo", strings.Join(ignore, "\n"))
}
