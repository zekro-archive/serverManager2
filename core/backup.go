package core

import (
	"os"
	"fmt"
	"strings"
	"time"
	"path/filepath"
	"github.com/zekroTJA/serverManager2/util"
)


type Backup struct {
	Name, Path string
	Date time.Time
}

func getBackups(screen *Screen, location string) *[]Backup {
	var out []Backup
	_, err := os.Stat(location)
	if os.IsNotExist(err) {
		os.MkdirAll(location, os.ModePerm)
	}
	filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
		folder := strings.Replace(path, location, "", -1)
		pstat, _ := os.Stat(path)
		if pstat.Mode().IsDir() && len(strings.Split(folder, "/")) == 2 {
			fmt.Println(path)
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

func BackupMenu(screen *Screen, config *util.Conf) {
	if config.BackupLocation == "" {
		util.LogError("Backup location is not set")
		util.Pause()
		return
	}

	for {
		backups := getBackups(screen, config.BackupLocation)
		
			fmt.Println("NAME               DATE")
			for _, e := range *backups {
				fmt.Println(
					columnize(e.Name, 17) + "  " + 
					e.Date.Format(time.RFC850))
			}
			_ = util.Cinpt("\n> ")
	}
}