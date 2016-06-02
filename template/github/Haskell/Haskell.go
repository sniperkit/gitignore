package Haskell

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"dist",
		"dist-*",
		"cabal-dev",
		"*.o",
		"*.hi",
		"*.chi",
		"*.chs.h",
		"*.dyn_o",
		"*.dyn_hi",
		".hpc",
		".hsenv",
		".cabal-sandbox/",
		"cabal.sandbox.config",
		"*.prof",
		"*.aux",
		"*.hp",
		"*.eventlog",
		".stack-work/",
		"cabal.project.local",
	}
	template.SetTemplate("GitHub/Haskell", strings.Join(ignore, "\n"))
}