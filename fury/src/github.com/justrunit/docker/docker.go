/*
	Fury service docker command handler
*/

package docker

import (
	"net/http"
	"os/exec"
	"log"
	"os"
	"io/ioutil"
	"strings"
	"github.com/justrunit/routeinit"
	"github.com/justrunit/furyutils"
	"github.com/justrunit/languages"
	"github.com/justrunit/furywebsockets"
)

func RunSnippetSync(resp http.ResponseWriter, req *http.Request) {
	validated, _, enc, body := routeinit.InitHandling(req, resp, []string{
		"language",
		"uid",
		"sid",
		"snippet"})
	if !validated {
		return
	}

	sidDetails := setContainerContext(body);
	results := executeContainer(sidDetails[ "uid" ].(string), sidDetails)
	enc.Encode(routeinit.ApiResponse{sidDetails[ "msg" ].(string), "", results, sidDetails[ "status" ].(int)})
	return
}

func RunSnippetAsync(resp http.ResponseWriter, req *http.Request) {
	validated, _, enc, body := routeinit.InitHandling(req, resp, []string{
		"language",
		"uid",
		"sid",
		"snippet"})
	if !validated {
		return
	}

	sid := body["sid"].(string)
	if furywebsockets.SidToOperation[ sid ] == "" {
		sidDetails := setContainerContext(body);
		executeContainerAsync(sidDetails)
		furywebsockets.SidToOperation[sid] = "run"
		enc.Encode(routeinit.ApiResponse{sidDetails[ "msg" ].(string), "", nil, sidDetails[ "status" ].(int)})
	} else {
		enc.Encode(routeinit.ApiResponse{"A " + furywebsockets.SidToOperation[sid] + " operation already running for current session", "", nil, 0})
	}
	return
}

/* Set container context like snippet directory for UID, deps install */
func setContainerContext(body map[string]interface{}) (sidDetails map[string]interface{}) {
	log.Println("Setting context for " + body["language"].(string) + " snippet UID " + body["uid"].(string) )

	status := 1
	msg := ""
	uid := body["uid"].(string)
	dir := "/home/justrunit/containers/" + uid

	// 1. Create snippet dir if it does not exist
	if !furyutils.DirExists(dir) {
		err := os.Mkdir(dir, 0777)
		exec.Command("bash", "-c", "chmod -R ugo+rw "+ dir).Output()
		if err != nil {
			msg = "Error creating container dir: " + err.Error()
			status = 0
			return
		}
	}

	// 2. Dump snippet into folder
	codeFilePath := dir + "/code"
	err := ioutil.WriteFile(codeFilePath, []byte(body["snippet"].(string)), 0777)
	if err != nil {
		msg = "Error writing to code file in container dir: " + err.Error()
	}

	// 3. Set language context
	setLanguageContext(dir, body["language"].(string))

	// Stash container details to create later on websocket initiation
	body[ "dir" ] = dir
	sidDetails = body
	sidDetails[ "msg" ] = msg
	sidDetails[ "uid" ] = uid
	sidDetails[ "status" ] = status

	status = 1
	return
}

/* Set language context like dependency files/folder initializations */
func setLanguageContext(dir string, language string) {

	// Read config file for languages
	languageConfigs := languages.GetLanguageConfigs()
	log.Println(languageConfigs)

	// Get code file path
	code := dir + "/code"

	// Language specific configs
	lc := languageConfigs[language].(map[string]interface{})
	cmd := "cat " + code + " | " + lc["deps_grep"].(string)

	// Generate deps
	deps, _ := exec.Command("bash", "-c", cmd).Output()

	// Write to deps file
	ioutil.WriteFile(dir + "/" + lc["deps_file"].(string), []byte(deps), 0777)
}

// Execute container, wait for it to complete, collect output and send
func executeContainer(uid string, sidDetails map[string]interface{}) (results []string) {
	dir := sidDetails["dir"].(string)
	output, err := exec.Command("bash", "-c", "docker run -v " +
		"\"" + dir + ":/home/justrunit/services/myproject\"" +
		" justrunit/" + sidDetails["language"].(string) ).Output()
	str := string(output) + err.Error()
	results = strings.Split(str, "\n")
	return
}

// Execute container, start it, wire up output to a channel
func executeContainerAsync(sidDetails map[string]interface{}) {
	dir := sidDetails["dir"].(string)
	dockerRunCmd := exec.Command("bash", "-c", "docker run -v " +
		"\"" + dir + ":/home/justrunit/services/myproject\"" +
		" justrunit/" + sidDetails["language"].(string) )
	dockerRunCmdOut, _ := dockerRunCmd.StdoutPipe()
	dockerRunCmdErr, _ := dockerRunCmd.StderrPipe()
	dockerRunCmd.Start()
	furywebsockets.ReaderToChannel(sidDetails[ "sid" ].(string),
		dockerRunCmdOut,
		dockerRunCmdErr);
}

