package CMake

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"CMakeCache.txt",
		"CMakeFiles",
		"CMakeScripts",
		"Makefile",
		"cmake_install.cmake",
		"install_manifest.txt",
		"CTestTestfile.cmake",
	}
	template.SetTemplate("GitHub/CMake", strings.Join(ignore, "\n"))
}