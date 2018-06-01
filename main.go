package main

import (
	"fmt"
	. "strconv"
	. "github.com/logrusorgru/aurora"
	"time"
	"github.com/zekroTJA/serverManager2/util"
	"github.com/zekroTJA/serverManager2/core"
)


const (
	VERSION = "2.3.0"
)

func getRunningSince(timestr string) string {
	started, _ := time.ParseInLocation("01/02/06 15:04:05", timestr, time.Now().Location())
	rsecs := time.Since(started).Seconds()
	days :=  int(rsecs / 86400)
	hours := int(rsecs) % 86400 / 3600
	mins :=  int(rsecs) % 86400 % 3600 / 60
	return fmt.Sprintf("[%03d:%02d:%02d]", days, hours, mins)
}


func printScreens(screens []core.Screen, servers []core.Screen, config util.Conf) string {
	util.Cls()
	fmt.Println(
		"Server Manager v." + VERSION,
		"\n(c) Ringo Hoffmann (zekro Development)",
		"\n\nServer Location: " + config.ServerLocation,
		"\nBackup Location: " + config.BackupLocation + "\n\n")
	for _, s := range servers {
		onof := Brown("[ STOPPED ]")
		if ok, sc := core.SliceContainsServer(screens, s); ok {
			onof = Green(getRunningSince(sc.Started))
		}
		fmt.Printf("%s %s %s\n", 
			Blue("[" + Itoa(s.Uid) + "]"), onof, s.Name)
		
	}
	return util.Cinpt("\n> ")
}

func main() {
	config := util.GetConf()

	var screens, servers []core.Screen

	res := ""
	for res != "exit" {
		screens = core.GetRunningScreens()
		servers = core.GetServers(config.ServerLocation)
		res = printScreens(screens, servers, config)
		core.HandleCmd(res, screens, servers, &config)
	}
}