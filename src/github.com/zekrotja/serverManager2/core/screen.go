package core

import (
	"os"
	"os/exec"
	_ "fmt"
	"strings"
	"regexp"
	"github.com/zekrotja/serverManager2/util"
	"path/filepath"
)


type Screen struct {
	Uid int
	Id, Name, Started string
}

func GetRunningScreens() []Screen {
	out, _ := exec.Command("screen", "-ls").Output()
	outsplit := strings.Split(string(out), "\n")
	regex := regexp.MustCompile(`[()]`)
	
	screens := []Screen {}
	for i, e := range outsplit[1:len(outsplit)-3]  {
		fields := strings.Fields(e)
		nameandid := strings.Split(fields[0], ".")
		screens = append(screens, Screen {
			i, 
			nameandid[0], 
			nameandid[1], 
			regex.ReplaceAllString(fields[1] + " " +  fields[2], ""),
		})
	}

	return screens
}

func GetServers(location string) []Screen {
	screens := []Screen {}
	filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
		folder := strings.Replace(path, location, "", -1)
		pathsplit := strings.Split(folder, "/")
		if len(pathsplit) == 2 {
			screens = append(screens, Screen {
				Uid: len(screens),
				Name: folder[1:] })
		}
		return err
	})
	return screens
}

func SliceContainsServer(slc []Screen, server Screen) bool {
	for _, e := range slc {
		if e.Name == server.Name {
			return true
		}
	}
	return false
}

// SCREEN ACTION FUNCTIONS

func StartScreen(screen Screen, screens []Screen, servers []Screen, config util.Conf) {

}

func StopScreen(screen Screen, screens []Screen, config util.Conf) {

}

func ResumeScreen(screen Screen, screens []Screen, config util.Conf) {

}