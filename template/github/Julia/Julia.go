package Julia

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*.jl.cov",
		"*.jl.*.cov",
		"*.jl.mem",
		"deps/deps.jl",
	}
	template.SetTemplate("GitHub/Julia", strings.Join(ignore, "\n"))
}