package CodeIgniter

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*/config/development",
		"*/logs/log-*.php",
		"!*/logs/index.html",
		"*/cache/*",
		"!*/cache/index.html",
		"!*/cache/.htaccess",
	}
	template.SetTemplate("GitHub/CodeIgniter", strings.Join(ignore, "\n"))
}