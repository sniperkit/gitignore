package tags

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Ignore tags created by etags, ctags, gtags (GNU global) and cscope",
		"TAGS",
		".TAGS",
		"!TAGS/",
		"tags",
		".tags",
		"!tags/",
		"gtags.files",
		"GTAGS",
		"GRTAGS",
		"GPATH",
		"cscope.files",
		"cscope.out",
		"cscope.in.out",
		"cscope.po.out",
		"",
	}
	template.Add("github/global/tags", strings.Join(ignore, "\n"))
}
