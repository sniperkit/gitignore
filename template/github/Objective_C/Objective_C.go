package Objective_C

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Xcode",
		"#",
		"# gitignore contributors: remember to update Global/Xcode.gitignore, Objective-C.gitignore & Swift.gitignore",
		"",
		"## Build generated",
		"build/",
		"DerivedData/",
		"",
		"## Various settings",
		"*.pbxuser",
		"!default.pbxuser",
		"*.mode1v3",
		"!default.mode1v3",
		"*.mode2v3",
		"!default.mode2v3",
		"*.perspectivev3",
		"!default.perspectivev3",
		"xcuserdata/",
		"",
		"## Other",
		"*.moved-aside",
		"*.xcuserstate",
		"",
		"## Obj-C/Swift specific",
		"*.hmap",
		"*.ipa",
		"*.dSYM.zip",
		"*.dSYM",
		"",
		"# CocoaPods",
		"#",
		"# We recommend against adding the Pods directory to your .gitignore. However",
		"# you should judge for yourself, the pros and cons are mentioned at:",
		"# https://guides.cocoapods.org/using/using-cocoapods.html#should-i-check-the-pods-directory-into-source-control",
		"#",
		"# Pods/",
		"",
		"# Carthage",
		"#",
		"# Add this line if you want to avoid checking in source code from Carthage dependencies.",
		"# Carthage/Checkouts",
		"",
		"Carthage/Build",
		"",
		"# fastlane",
		"#",
		"# It is recommended to not store the screenshots in the git repo. Instead, use fastlane to re-generate the",
		"# screenshots whenever they are needed.",
		"# For more information about the recommended setup visit:",
		"# https://github.com/fastlane/fastlane/blob/master/fastlane/docs/Gitignore.md",
		"",
		"fastlane/report.xml",
		"fastlane/screenshots",
		"",
		"#Code Injection",
		"#",
		"# After new code Injection tools there's a generated folder /iOSInjectionProject",
		"# https://github.com/johnno1962/injectionforxcode",
		"",
		"iOSInjectionProject/",
	}
	template.SetTemplate("GitHub/Objective_C", strings.Join(ignore, "\n"))
}