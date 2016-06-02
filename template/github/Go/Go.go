package Go

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Compiled Object files, Static and Dynamic libs (Shared Objects)",
		"*.o",
		"*.a",
		"*.so",
		"",
		"# Folders",
		"_obj",
		"_test",
		"",
		"# Architecture specific extensions/prefixes",
		"*.[568vq]",
		"[568vq].out",
		"",
		"*.cgo1.go",
		"*.cgo2.c",
		"_cgo_defun.c",
		"_cgo_gotypes.go",
		"_cgo_export.*",
		"",
		"_testmain.go",
		"",
		"*.exe",
		"*.test",
		"*.prof",
	}
	template.SetTemplate("GitHub/Go", strings.Join(ignore, "\n"))
}
