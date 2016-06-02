package SublimeText

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# cache files for sublime text",
		"*.tmlanguage.cache",
		"*.tmPreferences.cache",
		"*.stTheme.cache",
		"",
		"# workspace files are user-specific",
		"*.sublime-workspace",
		"",
		"# project files should be checked into the repository, unless a significant",
		"# proportion of contributors will probably not be using SublimeText",
		"# *.sublime-project",
		"",
		"# sftp configuration file",
		"sftp-config.json",
		"",
		"# Package control specific files",
		"Package Control.last-run",
		"Package Control.ca-list",
		"Package Control.ca-bundle",
		"Package Control.system-ca-bundle",
		"Package Control.cache/",
		"Package Control.ca-certs/",
		"bh_unicode_properties.cache",
		"",
		"# Sublime-github package stores a github token in this file",
		"# https://packagecontrol.io/packages/sublime-github",
		"GitHub.sublime-settings",
	}
	template.SetTemplate("GitHub/Global/SublimeText", strings.Join(ignore, "\n"))
}
