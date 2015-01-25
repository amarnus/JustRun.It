/*
	Languages meta data
*/

package languages

import (
	"encoding/json"
	"fmt"
)

func GetLanguageConfigs() (languages map[string]interface{}) {
	err := json.Unmarshal([]byte(`
		{
			"python": {
				"deps_grep": "grep -P '\\s*(?:from|import)' | perl -nle 'print $1 if /(?:import|from)\\s*([\\w\\-]+)/'",
				"deps_file": "requirements.txt",
				"install_deps": "virtualenv env && if [ -s \"requirements.txt\" ]; then pip install -r requirements.txt; fi;",
				"lint_error_regexes": [
					"invalid syntax"
				]
			},
			"ruby": {
				"deps_prefix": [
					"source 'https://rubygems.org'"
				],
				"deps_grep": "grep -P '\\s*(?:require )' | perl -nle 'print \"gem \\\"$1\\\"\" if /(?:require)\\s*.([\\w\\-]+)./'",
				"deps_file": "Gemfile",
				"install_deps": "if [ -s \"Gemfile\" ]; then bundle install; fi;",
				"lint_error_regexes": [
					"error",
					"undefined method"
				]
			},
			"nodejs": {
				"deps_grep": "grep -P 'require\\s*\\(' | perl -nle 'print $1 if /require\\s*\\(.(.*?).\\s*\\)/'",
				"deps_file": "deps",
				"install_deps": "if [ -s \"deps\" ]; then cat deps | xargs -l npm install; fi;",
				"lint_error_regexes": [
					"error",
					"undefined method"
				]
			},
			"php": {
				"deps_grep": "nodepsgrep",
				"deps_file": "composer.json",
				"install_deps": "if [ -s \"composer.json\" ]; then composer install --verbose; fi;",
				"lint_error_regexes": [
					"Errors parsing"
				]
			},
			"go": {
				"deps_grep": "sed ':a;N;$!ba;s/\n/ /g' | perl -nle 'if ($_ =~ /import\\s*\\(\\s*(.*?)\\s*\\)/) { $v = $1; $v =~ s/\\s+/\\n/g; $v =~ s/\\\"//g; print $v;}'",
				"deps_file": "Goopfile",
				"code_file": "code.go",
				"install_deps": "if [ -s \"Goopfile\" ]; then goop install; fi;",
				"lint_error_regexes": [
					"Errors"
				]
			}
		}
	`), &languages)
	if err != nil {
		fmt.Println(err)
	}
	return
}


