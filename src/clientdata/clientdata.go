/*
Copyright 2020 The Worker-Monitor-Client Author.
Licensed under the GNU General Public License v3.0.
    https://github.com/MfsTeller/worker-monitor-client/blob/master/LICENSE
*/
package clientdata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// ClientDataInterface defines interface for clientdata.
type ClientDataInterface interface {
	GetClientData() []ClientData
	Write(string, os.FileMode)
	Read(string)
	Get(int64) []byte
	Post([]byte)
}

// ClientData defines user identification data for worker-monitor-client.
type ClientData struct {
	ClientID         int64  `json:"client_id"`
	Name             string `json:"name"`
	ExecDatetime     string `json:"exec_datetime"`
	StartupDatetime  string `json:"startup_datetime"`
	ShutdownDatetime string `json:"shutdown_datetime"`
}

// ClientDataList defines multiple dates ClientData.
type ClientDataList struct {
	ClientData []ClientData
}

// NewClientData creates a client data controller.
func NewClientData(clientData []ClientData) *ClientDataList {
	w := new(ClientDataList)
	w.ClientData = clientData
	return w
}

// GetClientData is a getter for ClientData.
func (u *ClientDataList) GetClientData() []ClientData {
	return u.ClientData
}

// Write function creates a client data file.
func (u *ClientDataList) Write(filepath string, perm os.FileMode) {
	jsonBytes, err := json.Marshal(u.ClientData)
	if err != nil {
		log.Fatal(err)
	}
	out := new(bytes.Buffer)
	json.Indent(out, jsonBytes, "", "    ")
	err = ioutil.WriteFile(filepath, out.Bytes(), perm)
	if err != nil {
		log.Fatal(err)
	}
}

// Read function fetches a client data file.
func (u *ClientDataList) Read(filepath string) {
	jsonBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonBytes, &u.ClientData)
	if err != nil {
		log.Fatal(err)
	}
}

// HttpRequast creates http request for worker-monitor-server.
func HttpRequest(method string, url string, header map[string]string, body []byte) []byte {
	// create http request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	// set header
	for key, value := range header {
		req.Header.Set(key, value)
	}

	// execute http request
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// return response body
	byteList, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return byteList
}

// Get function creates a http GET request.
func (u *ClientDataList) Get(clientID int64) []byte {
	method := "GET"
	url := fmt.Sprintf(
		"http://192.168.99.100:8080/clientdata/%d",
		clientID,
	)
	respBody := HttpRequest(method, url, nil, nil)
	return respBody
}

// Get function creates a http POST request.
func (u *ClientDataList) Post(body []byte) {
	method := "POST"
	url := "http://192.168.99.100:8080/clientdata"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	HttpRequest(method, url, header, body)
}
