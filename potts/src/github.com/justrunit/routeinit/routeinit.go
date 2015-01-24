/*
	Fury service router
*/

package routeinit

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"net/url"
)

type ApiResponse struct {
	ErrorMessage string      `json:"error,omitempty"`
	Id           string      `json:"id,omitempty"`
	Result       interface{} `json:"result,omitempty"`
	Status       bool        `json:"status"`
}

/* Initializes variables needed to handle every request and authenticates and authorizes the request */
func InitHandling(
	req *http.Request,
	resp http.ResponseWriter,
	mandatoryBodyParams []string) (
	validated bool,
	urlParams map[string]string,
	enc *json.Encoder,
	body map[string]interface{}) {

	urlParams = mux.Vars(req)
	params := req.URL.Query()
	enc = json.NewEncoder(resp)

	validated, body = validateRequest(urlParams, params, req.Body, resp, enc)

	if validated {
		// Check for missing fields
		fields := ""
		for _, field := range mandatoryBodyParams {
			if body[field] == nil {
				if fields != "" {
					fields = fields + ","
				}
				fields = fields + field
				validated = false
			}
		}

		// If missing fields, set response
		if fields != "" {
			resp.WriteHeader(http.StatusBadRequest)
			enc.Encode(ApiResponse{fields + " not present", "", nil, false})
		}
	}

	return
}

/* Authenticate request and authorize action */
func validateRequest(
	urlVars map[string]string,
	params url.Values,
	body io.ReadCloser,
	resp http.ResponseWriter,
	enc *json.Encoder) (ok bool,
	dat map[string]interface{}) {

	/* read body and parse into interface{} */
	buf := new(bytes.Buffer)
	n, _ := buf.ReadFrom(body)
	ok = true
	if n != 0 {
		if err := json.Unmarshal(buf.Bytes(), &dat); err != nil {
			log.Println(err)
			resp.WriteHeader(http.StatusBadRequest)
			err = enc.Encode(ApiResponse{"malformed json", "", nil, false})
			ok = false
		}
	}

	return
}
