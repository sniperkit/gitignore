package Qt

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# C++ objects and libs",
		"",
		"*.slo",
		"*.lo",
		"*.o",
		"*.a",
		"*.la",
		"*.lai",
		"*.so",
		"*.dll",
		"*.dylib",
		"",
		"# Qt-es",
		"",
		"/.qmake.cache",
		"/.qmake.stash",
		"*.pro.user",
		"*.pro.user.*",
		"*.qbs.user",
		"*.qbs.user.*",
		"*.moc",
		"moc_*.cpp",
		"qrc_*.cpp",
		"ui_*.h",
		"Makefile",
		"*build-*",
		"",
		"# QtCreator",
		"",
		"*.autosave",
		"",
		"# QtCtreator Qml",
		"*.qmlproject.user",
		"*.qmlproject.user.*",
		"",
		"# QtCtreator CMake",
		"CMakeLists.txt.user",
		"",
	}
	template.SetTemplate("GitHub/Qt", strings.Join(ignore, "\n"))
}