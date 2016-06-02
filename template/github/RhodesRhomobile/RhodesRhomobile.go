package RhodesRhomobile

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"rholog-*",
		"sim-*",
		"bin/libs",
		"bin/RhoBundle",
		"bin/tmp",
		"bin/target",
		"bin/*.ap_",
		"*.o",
		"*.jar",
	}
	template.SetTemplate("GitHub/RhodesRhomobile", strings.Join(ignore, "\n"))
}