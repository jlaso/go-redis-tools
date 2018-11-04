package common

import (
	"os"
	"encoding/json"
	"sync"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"fmt"
)

const DEFAULTURL = "localhost:3000"

// @Todo: change configuration to accept more than one server
// @Todo: change configuration to accept servers with credentilas
type Configuration struct {
	Url  string `json:"url"`
	Port int    `json:"port"`
	Host string `json:"host"`
}

// http://marcio.io/2015/07/singleton-pattern-in-go/
var instance *Configuration
var once sync.Once

func configFileName() string {
	const configFolder = ".jlaso-redis-tools"
	const configFile = "config.json"
	usr, _ := user.Current()
	dir := usr.HomeDir
	os.Mkdir(filepath.Join(dir, configFolder), 0744)
	return filepath.Join(dir, configFolder, configFile)
}

func GetConfiguration() *Configuration {
	instance = &Configuration{
		Url:  DEFAULTURL,
		Port: 6379,
		Host: "localhost",
	}
	once.Do(func() {
		raw, err := ioutil.ReadFile(configFileName())
		if err == nil {
			json.Unmarshal(raw, &instance)
		}
		fmt.Println(instance)
	})
	return instance
}

func (cfg *Configuration) Save() error {
	data, err := json.Marshal(cfg)
	if err == nil {
		err = ioutil.WriteFile(configFileName(), data, 0744)
	}
	return err
}
