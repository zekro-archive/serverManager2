package main

import (
	"fmt"
	"os"
	. "strconv"
	. "github.com/logrusorgru/aurora"
	"time"
	"github.com/zekroTJA/serverManager2/util"
	"github.com/zekroTJA/serverManager2/core"
)


const (
	VERSION = "2.8.1"
)

func getRunningSince(timestr string) string {
	started, _ := time.ParseInLocation("01/02/06 15:04:05", timestr, time.Now().Location())
	rsecs := time.Since(started).Seconds()
	days :=  int(rsecs / 86400)
	hours := int(rsecs) % 86400 / 3600
	mins :=  int(rsecs) % 86400 % 3600 / 60
	return fmt.Sprintf("[%03d:%02d:%02d]", days, hours, mins)
}

func initLoggingPath() {
	path := "/etc/servermanager/logs"
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
}

func printScreens(screens *[]core.Screen, servers *[]core.Screen, config *util.Conf) string {
	util.Cls()
	fmt.Println(
		"Server Manager v." + VERSION,
		"\n(c) Ringo Hoffmann (zekro Development)",
		"\n\nServer Location: " + config.ServerLocation,
		"\nBackup Location: " + config.BackupLocation + "\n\n")
	for _, s := range *servers {
		onof := Brown("[ STOPPED ]")
		if ok, sc := core.SliceContainsServer(screens, &s); ok {
			onof = Green(getRunningSince(sc.Started))
		}
		index := func()string {
			if s.Uid < 10 {
				return "0" + Itoa(s.Uid)
			}
			return Itoa(s.Uid)
		}()
		fmt.Printf("%s %s %s\n", 
			Blue("[" + index + "]"), onof, s.Name)
		
	}
	return util.Cinpt("\n> ")
}

func main() {

	args := &core.Args {}
	testing := false
	for _, e := range os.Args {
		if e == "--test" {
			testing = true
		}
	}

	var config *util.Conf

	if testing {
		config = util.GetConf("./testconf.json")
	} else {
		config = util.GetConf()
	}

	initLoggingPath()

	var screens, servers *[]core.Screen

	res := ""
	for res != "exit" && res != "e" {
		screens = core.GetRunningScreens()
		servers = core.GetServers(config.ServerLocation)
		if args.Init(servers, screens, config) {
			args.Parse(VERSION)
			return
		}
		res = printScreens(screens, servers, config)
		core.HandleCmd(res, screens, servers, config)
	}
}