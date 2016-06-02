package Lua

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Compiled Lua sources",
		"luac.out",
		"",
		"# luarocks build files",
		"*.src.rock",
		"gitignore-master.zip",
		"*.tar.gz",
		"",
		"# Object files",
		"*.o",
		"*.os",
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
		"*.def",
		"*.exp",
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
	}
	template.SetTemplate("GitHub/Lua", strings.Join(ignore, "\n"))
}