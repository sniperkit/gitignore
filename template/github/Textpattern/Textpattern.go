package Textpattern

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		".htaccess",
		"css.php",
		"rpc/",
		"sites/site*/admin/",
		"sites/site*/private/",
		"sites/site*/public/admin/",
		"sites/site*/public/setup/",
		"sites/site*/public/theme/",
		"textpattern/",
		"HISTORY.txt",
		"README.txt",
	}
	template.SetTemplate("GitHub/Textpattern", strings.Join(ignore, "\n"))
}