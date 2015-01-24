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
)

var DockerContainers = make(map[string]interface{})

func RunSnippet(resp http.ResponseWriter, req *http.Request) {
	validated, _, enc, body := routeinit.InitHandling(req, resp, []string{
		"language",
		"uid",
		"snippet",
		"deps"})
	if !validated {
		return
	}

	status, msg, uid := setContainerContext(body);
	results := executeContainer(uid)
	enc.Encode(routeinit.ApiResponse{msg, "", results, status})
	return
}

/* Set container context like snippet directory for UID, deps install */
func setContainerContext(body map[string]interface{}) (status int, msg string, uid string) {
	log.Println("Setting context for " + body["language"].(string) + " snippet UID " + body["uid"].(string) )

	msg = ""
	uid = body["uid"].(string)
	dir := "/home/justrunit/containers/" + uid

	// 1. Create snippet dir if it does not exist
	if !furyutils.DirExists(dir) {
		err := os.MkdirAll(dir, 00777)
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
	DockerContainers[ uid ] = body

	status = 1
	return
}

/* Set language context like dependency files/folder initializations */
func setLanguageContext(dir string, language string) {
	if language == "python" {
		ioutil.WriteFile(dir + "/requirements.txt", []byte(""), 0700)
	}
}

func executeContainer(uid string) (results []string) {
	containerDetails := DockerContainers[ uid ].(map[string]interface{})
	dir := containerDetails["dir"].(string)
	// cmd := []string{
	output, _ := exec.Command("bash", "-c", "docker run -v " +
		"\"" + dir + ":/home/justrunit/services/myproject\"" +
		" justrunit/" + containerDetails["language"].(string) ).Output()
	str := string(output)
	results = strings.Split(str, "\n")
	return
}

