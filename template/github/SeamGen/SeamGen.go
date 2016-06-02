package SeamGen

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"/bootstrap/data",
		"/bootstrap/tmp",
		"/classes/",
		"/dist/",
		"/exploded-archives/",
		"/test-build/",
		"/test-output/",
		"/test-report/",
		"/target/",
		"temp-testng-customsuite.xml",
		"",
		"# based on http://stackoverflow.com/a/8865858/422476 I am removing inline comments",
		"",
		"#/classes/ all class files",
		"#/dist/ contains generated war files for deployment",
		"#/exploded-archives/ war content generation during deploy (or explode)",
		"#/test-build/ test compilation (ant target for Seam)",
		"#/test-output/ test results",
		"#/test-report/ test report generation for, e.g., Hudson",
		"#/target/ maven output folder",
		"#temp-testng-customsuite.xml generated when running test cases under Eclipse",
		"",
		"# Thanks to @VonC and @kraftan for their helpful answers on a related question",
		"# on StackOverflow.com:",
		"# http://stackoverflow.com/questions/4176687",
		"# /what-is-the-recommended-source-control-ignore-pattern-for-seam-projects",
	}
	template.SetTemplate("GitHub/SeamGen", strings.Join(ignore, "\n"))
}
