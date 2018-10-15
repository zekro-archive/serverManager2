package core

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	. "strings"

	"../util"
)

var AUTOSTARTFILE = "/etc/init.d/servermanager_scedule"

const (
	ERR_NULL = iota
	ERR_PATH_NOT_EXIST
	ERR_FILE_NOT_EXIST
	ERR_READ
)

func GetAutostart() (string, int) {
	pathsplit := Split(AUTOSTARTFILE, "/")
	path := Join(pathsplit[0:len(pathsplit)-1], "/")
	_, err := os.Stat(path)
	if err != nil {
		return "", ERR_PATH_NOT_EXIST
	}
	_, err = os.Stat(AUTOSTARTFILE)
	if err != nil {
		return "", ERR_FILE_NOT_EXIST
	}
	var datab []byte
	datab, err = ioutil.ReadFile(AUTOSTARTFILE)
	if err != nil {
		return "", ERR_READ
	}
	data := string(datab)
	out := Split(Split(data, "# SERVERS: ")[1], "\n")[0]
	return out, ERR_NULL
}

func CreateAutostart(servers *[]Screen) error {
	if _, stat := GetAutostart(); stat == ERR_PATH_NOT_EXIST {
		util.LogError("Path '/etc/init.d' does not exist!")
		util.Pause()
		return errors.New("ERR_PATH_NOT_EXIST")
	}
	var _servers []string
	for _, s := range *servers {
		_servers = append(_servers, s.Name)
	}
	out := fmt.Sprintf(
		"#!/bin/bash\n\n"+
			"# SERVERS: %s\n\n"+
			"%s -s %s --loop", Join(_servers, ","), os.Args[0], Join(_servers, ","))
	err := ioutil.WriteFile(AUTOSTARTFILE, []byte(out), 0777)
	return err
}

func ResetAutostart() error {
	return os.Remove(AUTOSTARTFILE)
}
