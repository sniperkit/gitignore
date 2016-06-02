package SlickEdit

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# SlickEdit workspace and project files are ignored by default because",
		"# typically they are considered to be developer-specific and not part of a",
		"# project.",
		"*.vpw",
		"*.vpj",
		"",
		"# SlickEdit workspace history and tag files always contain user-specific",
		"# data so they should not be stored in a repository.",
		"*.vpwhistu",
		"*.vpwhist",
		"*.vtg",
	}
	template.SetTemplate("GitHub/Global/SlickEdit", strings.Join(ignore, "\n"))
}
