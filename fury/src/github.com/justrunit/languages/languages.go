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
				"deps_grep": "grep -P '\\s*(?:from|import)' | perl -nle 'print $1 if /(?:import|from)\\s*(\\w+)/'",
				"deps_file": "requirements.txt"
			}
		}
	`), &languages)
	if err != nil {
		fmt.Println(err)
	}
	return
}


