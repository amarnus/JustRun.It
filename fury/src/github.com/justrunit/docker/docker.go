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
	"regexp"
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

	// 1. Set the container context
	sidDetails := setContainerContext(body);
	sidDetails[ "op" ] = "run"

	// 2. Run the container
	status := sidDetails[ "status" ].(int)
	var results []string
	var fstatus int = status
	if status == 1 {
		results, fstatus = executeContainer(sidDetails[ "uid" ].(string), sidDetails, 0)
	}

	enc.Encode(routeinit.ApiResponse{sidDetails[ "msg" ].(string), "", results, fstatus})
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
		sidDetails[ "op" ] = "run"
		executeContainerAsync(sidDetails)
		furywebsockets.SidToOperation[sid] = "run"
		enc.Encode(routeinit.ApiResponse{sidDetails[ "msg" ].(string), "", nil, sidDetails[ "status" ].(int)})
	} else {
		enc.Encode(routeinit.ApiResponse{"A " + furywebsockets.SidToOperation[sid] + " operation already running for current session", "", nil, 0})
	}
	return
}

func InstallDepsAsync(resp http.ResponseWriter, req *http.Request) {
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
		sidDetails[ "op" ] = "install-deps"
		executeContainerAsync(sidDetails)
		furywebsockets.SidToOperation[sid] = "install-deps"
		enc.Encode(routeinit.ApiResponse{sidDetails[ "msg" ].(string), "", nil, sidDetails[ "status" ].(int)})
	} else {
		enc.Encode(routeinit.ApiResponse{"A " + furywebsockets.SidToOperation[sid] + " operation already running for current session", "", nil, 0})
	}
	return
}

func InstallDepsSync(resp http.ResponseWriter, req *http.Request) {
	validated, _, enc, body := routeinit.InitHandling(req, resp, []string{
		"language",
		"uid",
		"sid",
		"snippet"})
	if !validated {
		return
	}

	// 1. Set the container context
	sidDetails := setContainerContext(body);
	sidDetails[ "op" ] = "install-deps"

	// 2. Run the container
	status := sidDetails[ "status" ].(int)
	var results []string
	var fstatus int = status
	if status == 1 {
		results, fstatus = executeContainer(sidDetails[ "uid" ].(string), sidDetails, 0)
	}

	enc.Encode(routeinit.ApiResponse{sidDetails[ "msg" ].(string), "", results, fstatus})
	return
}

func LintSnippetSync(resp http.ResponseWriter, req *http.Request) {
	validated, _, enc, body := routeinit.InitHandling(req, resp, []string{
		"language",
		"uid",
		"sid",
		"snippet"})
	if !validated {
		return
	}

	// 1. Set the container context
	sidDetails := setContainerContext(body);
	sidDetails[ "op" ] = "lint"

	// 2. Run the container
	status := sidDetails[ "status" ].(int)
	var results []string
	var fstatus int = status
	if status == 1 {
		results, fstatus = executeContainer(sidDetails[ "uid" ].(string), sidDetails, 1)
	}

	enc.Encode(routeinit.ApiResponse{sidDetails[ "msg" ].(string), "", results, fstatus})
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
	lc := setLanguageContext(dir, body["language"].(string), sidDetails[ "deps" ])

	// Stash container details to create later on websocket initiation
	body[ "dir" ] = dir
	sidDetails = body
	sidDetails[ "msg" ] = msg
	sidDetails[ "uid" ] = uid
	sidDetails[ "install_deps" ] = lc[ "install_deps" ].(string)
	sidDetails[ "status" ] = status

	status = 1
	return
}

/* Set language context like dependency files/folder initializations */
func setLanguageContext(dir string, language string, depsInput interface{}) (lc map[string]interface{}) {

	// Read config file for languages
	languageConfigs := languages.GetLanguageConfigs()
	log.Println(languageConfigs)

	// Get code file path
	code := dir + "/code"

	// Language specific configs
	lc = languageConfigs[language].(map[string]interface{})

	// Generate deps
	deps := []byte("")
	if depsInput == nil {
		cmd := "cat " + code + " | " + lc["deps_grep"].(string) + " | " + "grep -v -P \"^\\s*$\""

		// Generate deps
		deps, _ = exec.Command("bash", "-c", cmd).Output()

		// Add any prefix
		if lc["deps_prefix"] != nil {
			lcdp := lc["deps_prefix"].([]interface{})
			prefix := ""
			for _, dp := range lcdp {
				prefix = prefix + dp.(string) + "\n"
			}
			deps = []byte(prefix + string(deps))
		}
	} else {
		deps = depsInput.([]byte)
	}

	// Write to deps file
	ioutil.WriteFile(dir + "/" + lc["deps_file"].(string), []byte(deps), 0777)
	return
}

// Execute container, wait for it to complete, collect output and send
func executeContainer(uid string, sidDetails map[string]interface{}, isLint int) (results []string, status int) {

	status = 1

	// Get language
	language := sidDetails["language"].(string)

	// Code dir
	dir := sidDetails["dir"].(string)

	// Docker cmd
	dockerCmd := "docker run -v " + "\"" + dir + ":/home/justrunit/services/myproject\"" +
		" justrunit/" + language
	if isLint == 1 {
		dockerCmd = dockerCmd + " /bin/bash -c '$LINT_CMD code'"
	}
	if sidDetails[ "op" ] == "install-deps" {
		dockerCmd = dockerCmd + " /bin/bash -c '" + sidDetails["install_deps"].(string) + "'"
	}

	// Docker run
	dockerRunCmd := exec.Command("bash", "-c", dockerCmd)
	dockerRunCmdOut, _ := dockerRunCmd.StdoutPipe()
	dockerRunCmdErr, _ := dockerRunCmd.StderrPipe()
	dockerRunCmd.Start()

	stdoutBytes, _ := ioutil.ReadAll(dockerRunCmdOut)
	stderrBytes, _ := ioutil.ReadAll(dockerRunCmdErr)

	str := string(stdoutBytes) + "\nStderr\n" + string(stderrBytes);

	// Check for lint errors
	if isLint == 1 {
		languageConfigs := languages.GetLanguageConfigs()
		lc := languageConfigs[ language ].(map[string]interface{})
		rs := lc[ "lint_error_regexes" ].([]interface{})
		for _, regex := range rs {
			log.Println("Checking for lint error: " + regex.(string))
			match, _ := regexp.MatchString(regex.(string), str)
			if match {
				status = 0
			}
		}
	}

	results = strings.Split(str, "\n")
	return
}

// Execute container, start it, wire up output to a channel
func executeContainerAsync(sidDetails map[string]interface{}) {

	// Parse language
	language := sidDetails[ "language" ].(string)

	// Code dir
	dir := sidDetails["dir"].(string)

	// Docker cmd
	dockerCmd := "docker run -v " +
		"\"" + dir + ":/home/justrunit/services/myproject\"" +
		" justrunit/" + language
	if sidDetails[ "op" ] == "install-deps" {
		dockerCmd = dockerCmd + " /bin/bash -c '" + sidDetails["install_deps"].(string) + "'"
	}
	dockerRunCmd := exec.Command("bash", "-c", dockerCmd )

	// Retrieve Output pipe handlers
	dockerRunCmdOut, _ := dockerRunCmd.StdoutPipe()
	dockerRunCmdErr, _ := dockerRunCmd.StderrPipe()

	// Start the snippet execution
	dockerRunCmd.Start()

	// Pipe the IO Readers to channels feeding the websockets
	furywebsockets.ReaderToChannel(sidDetails[ "sid" ].(string),
		dockerRunCmdOut,
		dockerRunCmdErr);
}

