package synopsysvcs

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Waveform formats",
		"*.vcd",
		"*.vpd",
		"*.evcd",
		"*.fsdb",
		"",
		"# Default name of the simulation executable.  A different name can be",
		"# specified with this switch (the associated daidir database name is",
		"# also taken from here):  -o <path>/<filename>",
		"simv",
		"",
		"# Generated for Verilog and VHDL top configs",
		"simv.daidir/",
		"simv.db.dir/",
		"",
		"# Infrastructure necessary to co-simulate SystemC models with",
		"# Verilog/VHDL models.  An alternate directory may be specified with this",
		"# switch:  -Mdir=<directory_path>",
		"csrc/",
		"",
		"# Log file - the following switch allows to specify the file that will be",
		"# used to write all messages from simulation:  -l <filename>",
		"*.log",
		"",
		"# Coverage results (generated with urg) and database location.  The",
		"# following switch can also be used:  urg -dir <coverage_directory>.vdb",
		"simv.vdb/",
		"urgReport/",
		"",
		"# DVE and UCLI related files.",
		"DVEfiles/",
		"ucli.key",
		"",
		"# When the design is elaborated for DirectC, the following file is created",
		"# with declarations for C/C++ functions.",
		"vc_hdrs.h",
		"",
	}
	template.Add("github/global/synopsysvcs", strings.Join(ignore, "\n"))
}
