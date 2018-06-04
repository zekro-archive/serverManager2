package util

import (
	"encoding/json"
	"os"
	"fmt"
	. "strings"
	"strconv"
	"io/ioutil"
)


var CONFFILE = "/etc/servermanager/config.json"


type Conf struct {
	ServerLocation string `json:"serverLocation"`
	BackupLocation string `json:"backupLocation"`
	Logging        int    `json:"enableLogging"`
}

func cutSlashAtBack(a string) string {
	if HasSuffix(a, "/") {
		a = a[0:len(a) - 1]
	}
	return a
}

func GetConf(loc ...string) *Conf {
	if len(loc) > 0 {
		CONFFILE = loc[0]
	}
	f, err := os.Open(CONFFILE)
	if os.IsNotExist(err) {
		newconf := Conf {}
		CreateConf(&newconf)
		fmt.Println("LOC: " + newconf.ServerLocation)
		return &newconf
	} else if err != nil {
		LogFatal("Failed reading config file:\n" + err.Error())
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	config := Conf {}
	decoder.Decode(&config)
	return &config
}

func CreateConf(current *Conf) {
	Cls()
	fmt.Printf("\nCONFIG EDITOR\n\nPlease only use total paths!\n\n")
	
	current.ServerLocation = cutSlashAtBack(Cinpt("serverLocation (current \"" + current.ServerLocation + "\") [Path]:\n> "))
	current.BackupLocation = cutSlashAtBack(Cinpt("backupLocation: (current \"" + current.BackupLocation + "\") [Path]:\n> "))
	current.Logging = func()int {
		for {
			inpt := Cinpt(fmt.Sprintf("enableLogging: (current \"%d\") [0/1]:\n> ", current.Logging))
			if inpt == "0" || inpt == "1" {
				res, err := strconv.Atoi(inpt)
				if err == nil {
					return res
				}
			}
			if inpt == "" {
				return 0
			}
			LogError("Enter a valid value for 'enableLogging' [0/1]!")
		}
	}()

	bjson, _ := json.MarshalIndent(current, "", "  ")

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
}