package main

import (
	"bufio"
    "fmt"
	"os"
	"os/exec"
	. "strconv"
	. "github.com/logrusorgru/aurora"
	"github.com/zekrotja/serverManager2/util"
	"github.com/zekrotja/serverManager2/core"
)


const (
	VERSION = "2.1.0"
)


func cinpt(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return text[0:len(text)-1]
}

func cls() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printScreens(screens []core.Screen, servers []core.Screen, config util.Conf) string {
	//cls()
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
	return cinpt("\n> ")
}

func main() {
	config := util.GetConf()

	var screens, servers []core.Screen

	res := ""
	for res != "exit" {
		screens = core.GetRunningScreens()
		servers = core.GetServers(config.ServerLocation)
		res = printScreens(screens, servers, config)
		core.HandleCmd(res, screens, servers, config)
	}
}