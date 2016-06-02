package Rails

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"*.rbc",
		"capybara-*.html",
		".rspec",
		"/log",
		"/tmp",
		"/db/*.sqlite3",
		"/db/*.sqlite3-journal",
		"/public/system",
		"/coverage/",
		"/spec/tmp",
		"**.orig",
		"rerun.txt",
		"pickle-email-*.html",
		"",
		"# TODO Comment out these rules if you are OK with secrets being uploaded to the repo",
		"config/initializers/secret_token.rb",
		"config/secrets.yml",
		"",
		"# dotenv",
		"# TODO Comment out this rule if environment variables can be committed",
		".env",
		"",
		"## Environment normalization:",
		"/.bundle",
		"/vendor/bundle",
		"",
		"# these should all be checked in to normalize the environment:",
		"# Gemfile.lock, .ruby-version, .ruby-gemset",
		"",
		"# unless supporting rvm < 1.11.0 or doing something fancy, ignore this:",
		".rvmrc",
		"",
		"# if using bower-rails ignore default bower_components path bower.json files",
		"/vendor/assets/bower_components",
		"*.bowerrc",
		"bower.json",
		"",
		"# Ignore pow environment settings",
		".powenv",
		"",
		"# Ignore Byebug command history file.",
		".byebug_history",
	}
	template.SetTemplate("GitHub/Rails", strings.Join(ignore, "\n"))
}
