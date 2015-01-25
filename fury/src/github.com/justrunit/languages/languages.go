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
				"lint_error_regexes": [
					"error",
					"undefined method"
				]
			}
		}
	`), &languages)
	if err != nil {
		fmt.Println(err)
	}
	return
}


