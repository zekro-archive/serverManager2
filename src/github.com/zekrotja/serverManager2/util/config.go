package util

import (
	"encoding/json"
	"os"
	_ "fmt"
	"io/ioutil"
)


const CONFFILE = "config.json"


type Conf struct {
	ServerLocation string `json:"serverLocation"`
	BackupLocation string `json:"backupLocation"`
}

func GetConf() Conf {
	f, err := os.Open(CONFFILE)
	if os.IsNotExist(err) {
		CreateConf()
		LogWarn("config file was not existent and was created now.\nPlease enter preferences in the config file and restart.")
		os.Exit(0)
	} else if err != nil {
		LogFatal("Failed reading config file:\n" + err.Error())
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	config := Conf {}
	decoder.Decode(&config)
	return config
}

func CreateConf() {
	stdConf := Conf {"", ""}
	bjson, _ := json.MarshalIndent(stdConf, "", "  ")
	err := ioutil.WriteFile(CONFFILE, bjson, 0644)
	if err != nil {
		LogFatal("Failed creating config file:\n" + err.Error())
	}
}