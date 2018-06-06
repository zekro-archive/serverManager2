package core

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"errors"
	"time"
	"path/filepath"
	"github.com/zekroTJA/serverManager2/util"
)


type Backup struct {
	Name, Path string
	Date time.Time
}

// PRIVATE FUNCTIONS

func getBackups(screen *Screen, location string) *[]Backup {
	var out []Backup
	_, err := os.Stat(location)
	if os.IsNotExist(err) {
		os.MkdirAll(location, os.ModePerm)
	}
	filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
		folder := strings.Replace(path, location + "/" + screen.Name, "", -1)
		pstat, _ := os.Stat(path)
		if pstat.Mode().IsDir() && len(strings.Split(folder, "/")) == 2 {
			out = append(out, Backup {
				folder,
				path,
				pstat.ModTime() })
		}
		return err
	})
	return &out
}

func columnize(cont string, csize int) string {
	if len(cont) > csize {
		return cont[0:csize-3] + "..."
	}
	return cont + "                              "[0:csize-len(cont)]
}

func fetchBackupByInd(inpt string, backups *[]Backup) (*Backup, error) {
	i, err := strconv.Atoi(inpt)
	if err != nil {
		return &Backup {}, err
	}
	if len(*backups) >= i {
		return &Backup {}, errors.New("out of bounds")
	}
	return &(*backups)[i], nil
}

// PUBLIC FUNCTIONS

func CreateBackup(screen *Screen, config *util.Conf, name string) {
	util.CopyDir(
		config.ServerLocation + "/" + screen.Name, 
	    config.BackupLocation + "/" + screen.Name + "/" + name)
}

func DeleteBackup(backup *Backup) {
	os.RemoveAll(backup.Path)
}

func BackupMenu(screen *Screen, config *util.Conf) {
	if config.BackupLocation == "" {
		util.LogError("Backup location is not set")
		util.Pause()
		return
	}

	inpt := ""
	for inpt != "exit" && inpt != "e" {
		backups := getBackups(screen, config.BackupLocation)
		
		fmt.Println("INDEX  " + columnize("NAME", 17) + "  DATE")
		for i, e := range *backups {
			fmt.Println(
				columnize(strconv.Itoa(i), 5) + "  " +
				columnize(e.Name, 17) + "  " + 
				e.Date.Format(time.RFC850))
		}
		inpt = util.Cinpt("\n> ")
		inptsplit := strings.Split(inpt, " ")
		invoke := inptsplit[0]
		args := inptsplit[1:]

		if len(args) > 0 {
			tscreen, err := fetchBackupByInd(args[0], backups)
			switch invoke {
			case "create":
				name := strings.Join(args, " ")
				CreateBackup(screen, config, name)
			case "delete":
				if err == nil {
					DeleteBackup(tscreen)
				}
			}
		}
	}
}