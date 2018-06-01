package util

import (
	"encoding/json"
	"os"
	"fmt"
	. "strings"
	"io/ioutil"
)


const CONFFILE = "/etc/servermanager/config.json"


type Conf struct {
	ServerLocation string `json:"serverLocation"`
	BackupLocation string `json:"backupLocation"`
}

func GetConf() Conf {
	f, err := os.Open(CONFFILE)
	if os.IsNotExist(err) {
		return CreateConf(Conf {"", ""})
	} else if err != nil {
		LogFatal("Failed reading config file:\n" + err.Error())
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	fmt.Println("CONFIG READ")
	config := Conf {}
	decoder.Decode(&config)
	return config
}

func CreateConf(current Conf) Conf {
	Cls()
	fmt.Printf("\nCONFIG EDITOR\n\nPlease only use total paths!\n\n")
	inptconf := Conf {
		Cinpt("serverLocation (current \"" + current.ServerLocation + "\"):\n> "), 
		Cinpt("backupLocation: (current \"" + current.BackupLocation + "\"):\n> ")}

	bjson, _ := json.MarshalIndent(inptconf, "", "  ")

	pathsplit := Split(CONFFILE, "/")
	path := Join(pathsplit[0:len(pathsplit)-1], "/")
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	err = ioutil.WriteFile(CONFFILE, bjson, 0644)
	if err != nil {
		LogFatal("Failed creating config file:\n" + err.Error())
	}
	return inptconf
}