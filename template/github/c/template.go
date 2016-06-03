package c

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Object files",
		"*.o",
		"*.ko",
		"*.obj",
		"*.elf",
		"",
		"# Precompiled Headers",
		"*.gch",
		"*.pch",
		"",
		"# Libraries",
		"*.lib",
		"*.a",
		"*.la",
		"*.lo",
		"",
		"# Shared objects (inc. Windows DLLs)",
		"*.dll",
		"*.so",
		"*.so.*",
		"*.dylib",
		"",
		"# Executables",
		"*.exe",
		"*.out",
		"*.app",
		"*.i*86",
		"*.x86_64",
		"*.hex",
		"",
		"# Debug files",
		"*.dSYM/",
		"*.su",
		"",
	}
	template.Add("github/c", strings.Join(ignore, "\n"))
}