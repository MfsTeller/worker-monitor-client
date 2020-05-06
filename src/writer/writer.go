package writer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Writer interface {
	GetClientData() []ClientData
	Write(string, os.FileMode)
	Get(int64) []byte
	Post([]byte)
}

type ClientData struct {
	ClientID         int64  `json:"client_id"`
	Name             string `json:"name"`
	ExecDatetime     string `json:"exec_datetime"`
	StartupDatetime  string `json:"startup_datetime"`
	ShutdownDatetime string `json:"shutdown_datetime"`
}

type WriterData struct {
	clientData []ClientData
}

func NewWriter(clientData []ClientData) *WriterData {
	w := new(WriterData)
	w.clientData = clientData
	return w
}

func (u *WriterData) GetClientData() []ClientData {
	return u.clientData
}

func (u *WriterData) Write(filepath string, perm os.FileMode) {
	jsonBytes, err := json.Marshal(u.clientData)
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

func (u *WriterData) Get(clientID int64) []byte {
	method := "GET"
	url := fmt.Sprintf(
		"http://192.168.99.100:8080/clientdata/%d",
		clientID,
	)
	respBody := HttpRequest(method, url, nil, nil)
	return respBody
}

func (u *WriterData) Post(body []byte) {
	method := "POST"
	url := "http://192.168.99.100:8080/clientdata"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	HttpRequest(method, url, header, body)
}
