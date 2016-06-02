package Android

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Built application files",
		"*.apk",
		"*.ap_",
		"",
		"# Files for the ART/Dalvik VM",
		"*.dex",
		"",
		"# Java class files",
		"*.class",
		"",
		"# Generated files",
		"bin/",
		"gen/",
		"out/",
		"",
		"# Gradle files",
		".gradle/",
		"build/",
		"",
		"# Local configuration file (sdk path, etc)",
		"local.properties",
		"",
		"# Proguard folder generated by Eclipse",
		"proguard/",
		"",
		"# Log Files",
		"*.log",
		"",
		"# Android Studio Navigation editor temp files",
		".navigation/",
		"",
		"# Android Studio captures folder",
		"captures/",
		"",
		"# Intellij",
		"*.iml",
		".idea/workspace.xml",
		"",
		"# Keystore files",
		"*.jks",
	}
	template.SetTemplate("GitHub/Android", strings.Join(ignore, "\n"))
}