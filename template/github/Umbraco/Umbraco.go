package Umbraco

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Note: VisualStudio gitignore rules may also be relevant",
		"",
		"# Umbraco",
		"# Ignore unimportant folders generated by Umbraco",
		"**/App_Data/Logs/",
		"**/App_Data/[Pp]review/",
		"**/App_Data/TEMP/",
		"**/App_Data/NuGetBackup/",
		"",
		"# Ignore Umbraco content cache file",
		"**/App_Data/umbraco.config",
		"",
		"# Don't ignore Umbraco packages (VisualStudio.gitignore mistakes this for a NuGet packages folder)",
		"# Make sure to include details from VisualStudio.gitignore BEFORE this",
		"!**/App_Data/[Pp]ackages/",
		"!**/[Uu]mbraco/[Dd]eveloper/[Pp]ackages",
		"",
		"# ImageProcessor DiskCache",
		"**/App_Data/cache/",
	}
	template.SetTemplate("GitHub/Umbraco", strings.Join(ignore, "\n"))
}