package core

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"../util"
)

type Backup struct {
	Name, Path string
	Date       time.Time
}

// PRIVATE FUNCTIONS

func columnize(cont string, csize int) string {
	if len(cont) > csize {
		return cont[0:csize-3] + "..."
	}
	return cont + "                              "[0:csize-len(cont)]
}

func fetchBackupByInd(inpt string, backups *[]Backup) (*Backup, error) {
	i, err := strconv.Atoi(inpt)
	if err != nil {
		return &Backup{}, err
	}
	if len(*backups) <= i {
		return &Backup{}, errors.New("out of bounds")
	}
	return &(*backups)[i], nil
}

// PUBLIC FUNCTIONS

func GetBackups(screen *Screen, location string) *[]Backup {
	var out []Backup
	_, err := os.Stat(location)
	if os.IsNotExist(err) {
		os.MkdirAll(location, os.ModePerm)
	}
	filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
		folder := strings.Replace(path, location+"/"+screen.Name, "", -1)
		pstat, _ := os.Stat(path)
		if pstat.Mode().IsDir() && len(strings.Split(folder, "/")) == 2 {
			out = append(out, Backup{
				folder,
				path,
				pstat.ModTime()})
		}
		return err
	})
	return &out
}

func CreateBackup(screen *Screen, config *util.Conf, name string) {
	util.CopyDir(
		config.ServerLocation+"/"+screen.Name,
		config.BackupLocation+"/"+screen.Name+"/"+name)
}

func DeleteBackup(backup *Backup) {
	os.RemoveAll(backup.Path)
}

func RevokeBackup(backup *Backup, config *util.Conf, path, name string, yes bool) {
	res := util.Cinpt(
		"ATTENTION:\nBefore restoring, current live state will be saved as backup here.\n" +
			"Also, the server should be shut down before restoring a backup!\n" +
			"Do you really want to restore? [yN]: ")
	if strings.ToLower(res) == "y" || yes {
		util.CopyDir(
			config.ServerLocation+"/"+name,
			config.BackupLocation+"/"+name+"/AUTO_"+time.Now().Format("02-01-06_15-04-05"))
		os.RemoveAll(path)
		util.CopyDir(
			backup.Path,
			path)
	}
}

func BackupMenu(screen *Screen, config *util.Conf) {
	if config.BackupLocation == "" {
		util.LogError("Backup location is not set")
		util.Pause()
		return
	}

	inpt := ""
	for inpt != "exit" && inpt != "e" {
		backups := GetBackups(screen, config.BackupLocation)

		util.Cls()
		fmt.Println(
			"BACKUP MANAGER - SCREEN: " + screen.Name + "\n\n" +
				"INDEX  " + columnize("NAME", 17) + "  DATE")
		for i, e := range *backups {
			fmt.Println(
				columnize(strconv.Itoa(i), 5) + "  " +
					columnize(e.Name, 17) + "  " +
					e.Date.Format(time.RFC850))
		}
		inpt = util.Cinpt("\nbackup > ")
		inptsplit := strings.Split(inpt, " ")
		invoke := inptsplit[0]
		args := inptsplit[1:]

		if invoke == "help" {
			fmt.Println(
				"\n create <name>   | Create a backup with the given name" +
					"\n delete <index>  | Delete backup by index" +
					"\n restore <index> | Restore a backup to live server\n")
			util.Cinpt("[Enter to continue...]")
			continue
		}

		if len(args) > 0 {
			tbackup, err := fetchBackupByInd(args[0], backups)
			switch invoke {
			case "create":
				name := strings.Join(args, " ")
				CreateBackup(screen, config, name)
			case "delete":
				if err == nil {
					DeleteBackup(tbackup)
				} else {
					util.LogError(err.Error())
					util.Cinpt("[Enter to continue...]")
				}
			case "restore":
				if err == nil {
					RevokeBackup(tbackup, config, config.ServerLocation+"/"+screen.Name, screen.Name, false)
				} else {
					util.LogError(err.Error())
					util.Cinpt("[Enter to continue...]")
				}
			}
		}
	}
}
