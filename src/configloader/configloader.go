/*
Copyright 2020 The Worker-Monitor-Client Author.
Licensed under the GNU General Public License v3.0.
    https://github.com/MfsTeller/worker-monitor-client/blob/master/LICENSE
*/
package configloader

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ConfigLoader is a interface for config data file.
type ConfigLoader interface {
	Load(string)
	GetWorkDir() string
	GetClientID() int64
	GetName() string
}

// ConfigData includes client information from config data file.
type ConfigData struct {
	ClientID int64  `json:"client_id"`
	Name     string `json:"name"`
	WorkDir  string `json:"work_dir"`
}

// NewConfigLoader creates a config loader.
func NewConfigLoader() *ConfigData {
	w := new(ConfigData)
	return w
}

// Load function fetches config file data.
func (u *ConfigData) Load(filepath string) {
	jsonBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	// decode
	err = json.Unmarshal(jsonBytes, u)
	if err != nil {
		log.Fatal(err)
	}
}

// GetWorkDir is a getter for working directory.
func (u *ConfigData) GetWorkDir() string {
	return u.WorkDir
}

// GetClientID is a getter for client ID.
func (u *ConfigData) GetClientID() int64 {
	return u.ClientID
}

// GetName is a getter for user name.
func (u *ConfigData) GetName() string {
	return u.Name
}
