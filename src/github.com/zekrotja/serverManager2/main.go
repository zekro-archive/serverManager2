package main

import (
    "fmt"
	. "strconv"
	. "github.com/logrusorgru/aurora"
	"github.com/zekrotja/serverManager2/util"
	"github.com/zekrotja/serverManager2/core"
)


const (
	VERSION = "2.1.0"
)

func printScreens(screens []core.Screen, servers []core.Screen, config util.Conf) string {
	util.Cls()
	fmt.Println(
		"Server Manager v." + VERSION,
		"\n(c) Ringo Hoffmann (zekro Development)",
		"\n\nServer Location: " + config.ServerLocation,
		"\nBackup Location: " + config.BackupLocation + "\n\n")
	for _, s := range servers {
		onof := Brown("[STOPPED]")
		if core.SliceContainsServer(screens, s) {
			onof = Green("[RUNNING]")
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