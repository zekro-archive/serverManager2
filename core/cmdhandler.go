package core

import (
	"fmt"
	. "strings"
	. "strconv"
	"github.com/zekroTJA/serverManager2/util"
)

func printHelp() {
	fmt.Println(
		"\n start <index/name> [e]   | Start server",
		"\n                          | Attach argument 'e' to run server in endless loop,",
		"\n                          | so it will restart after crash",
		"\n stop <index/name>        | Stop server",
		"\n resume <index/name>      | Resume a running server",
		"\n restart <index/name> [e] | Restart a server",
		"\n                          | Use 'e' as argument as same as with the start cmd",
		"\n backup <index/name>      | Start backup manager for specific server",
		"\n config (<editor>)        | Edit config of the program",
		"\n                          | If you pass the editor, you can edit the config in the editor",
		"\n                          | Example: 'config nano'",
		"\n exit                     | Exit the program",
		"\n\nConfig File Location: " + util.CONFFILE + "\n")
	util.Pause()
}

func fetchServer(servers *[]Screen, invoke string) *Screen {
	invoke = ToLower(invoke)
	for _, e := range *servers {
		invokei, err := Atoi(invoke)
		if err == nil && e.Uid == invokei {
			return &e
		} else if ToLower(e.Name) == invoke {
			return &e
		}
	}
	for _, e := range *servers {
		if HasPrefix(ToLower(e.Name), invoke) {
			return &e
		}
	}
	return &Screen {}
}

func HandleCmd(cmd string, screens *[]Screen, servers *[]Screen, config *util.Conf ) {
	cmdsplit := Split(cmd, " ")
	invoke := cmdsplit[0]
	args := cmdsplit[1:]

	switch len(args) {

	case 0:
		switch invoke {
		case "help":
			printHelp()
		case "config":
			util.CreateConf(config)
		}

	default:
		server := fetchServer(servers, args[0])
		if server == (&Screen {}) {
			util.LogError("Can not fetch '" + args[0] + "' to any server")
			return
		}
		switch invoke {
		case "start":
			endless := (len(args) > 1 && args[1] == "e")
			StartScreen(server, screens, config, endless)
		case "stop":
			StopScreen(server, screens, config)
		case "resume":
			ResumeScreen(server, screens, config)
		case "restart":
			endless := (len(args) > 1 && args[1] == "e")
			RestartScreen(server, screens, config, endless)
		case "backup":
			BackupMenu(server, config)
		case "config":
			util.EditConfWithEditor(config, args[0])
		}
	}
}