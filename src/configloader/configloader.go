package configloader

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ConfigLoader interface {
	Load(string)
	GetWorkDir() string
	GetClientID() int64
	GetName() string
}

type ConfigData struct {
	ClientID int64  `json:"client_id"`
	Name     string `json:"name"`
	WorkDir  string `json:"work_dir"`
}

func NewConfigLoader() *ConfigData {
	w := new(ConfigData)
	return w
}

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

func (u *ConfigData) GetWorkDir() string {
	return u.WorkDir
}

func (u *ConfigData) GetClientID() int64 {
	return u.ClientID
}

func (u *ConfigData) GetName() string {
	return u.Name
}
