package matlab

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"##---------------------------------------------------",
		"## Remove autosaves generated by the Matlab editor",
		"## We have git for backups!",
		"##---------------------------------------------------",
		"",
		"# Windows default autosave extension",
		"*.asv",
		"",
		"# OSX / *nix default autosave extension",
		"*.m~",
		"",
		"# Compiled MEX binaries (all platforms)",
		"*.mex*",
		"",
		"# Simulink Code Generation",
		"slprj/",
		"",
		"# Session info",
		"octave-workspace",
		"",
	}
	template.Add("github/global/matlab", strings.Join(ignore, "\n"))
}