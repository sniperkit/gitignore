package Jekyll

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"_site/",
		".sass-cache/",
		".jekyll-metadata",
	}
	template.SetTemplate("GitHub/Jekyll", strings.Join(ignore, "\n"))
}