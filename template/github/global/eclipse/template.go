package eclipse

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".metadata",
		"bin/",
		"tmp/",
		"*.tmp",
		"*.bak",
		"*.swp",
		"*~.nib",
		"local.properties",
		".settings/",
		".loadpath",
		".recommenders",
		"",
		"# Eclipse Core",
		".project",
		"",
		"# External tool builders",
		".externalToolBuilders/",
		"",
		"# Locally stored \"Eclipse launch configurations\"",
		"*.launch",
		"",
		"# PyDev specific (Python IDE for Eclipse)",
		"*.pydevproject",
		"",
		"# CDT-specific (C/C++ Development Tooling)",
		".cproject",
		"",
		"# JDT-specific (Eclipse Java Development Tools)",
		".classpath",
		"",
		"# Java annotation processor (APT)",
		".factorypath",
		"",
		"# PDT-specific (PHP Development Tools)",
		".buildpath",
		"",
		"# sbteclipse plugin",
		".target",
		"",
		"# Tern plugin",
		".tern-project",
		"",
		"# TeXlipse plugin",
		".texlipse",
		"",
		"# STS (Spring Tool Suite)",
		".springBeans",
		"",
		"# Code Recommenders",
		".recommenders/",
		"",
	}
	template.Add("github/global/eclipse", strings.Join(ignore, "\n"))
}
