/*
	Utils
*/

package furyutils

import (
	"os"
	"log"
)

func DirExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil { log.Println(path + " exists"); return true }
	if os.IsNotExist(err) { log.Println(path + " does not exist"); return false }
	return false
}

