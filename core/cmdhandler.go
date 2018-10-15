package core

import (
	"fmt"
	. "strconv"
	. "strings"

	"../util"
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
		"\n\nConfig File Location: "+util.CONFFILE+"\n")
	util.Pause()
}

func fetchServer(servers *[]Screen, invoke string) (*Screen, bool) {
	invoke = ToLower(invoke)
	for _, e := range *servers {
		invokei, err := Atoi(invoke)
		if err == nil && e.Uid == invokei {
			return &e, true
		} else if ToLower(e.Name) == invoke {
			return &e, true
		}
	}
	for _, e := range *servers {
		if HasPrefix(ToLower(e.Name), invoke) {
			return &e, true
		}
	}
	return &Screen{}, false
}

func testIfFetchSuccessfull(invoke string, ok bool, cb func()) {
	if !ok {
		util.LogError("Can not fetch '" + invoke + "' to any server")
		util.Pause()
	} else {
		cb()
	}
}

func HandleCmd(cmd string, screens *[]Screen, servers *[]Screen, config *util.Conf) {
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
		case "autostart":
			aservers, stat := GetAutostart()
			switch stat {
			case ERR_PATH_NOT_EXIST:
				util.LogError("Autostart path '/etc/init.d' can not be found!")
			case ERR_FILE_NOT_EXIST:
				util.LogWarn("Autostart entry not created yet.")
			case ERR_READ:
				util.LogError("An unexpected error occured while reading file.")
			default:
				util.LogInfo("Current servers in autostart: " + aservers)
			}
			util.Pause()
		}

	default:
		server, ok := fetchServer(servers, args[0])
		// if !ok {
		// 	util.LogError("Can not fetch '" + args[0] + "' to any server")
		// 	return
		// }
		switch invoke {
		case "start":
			testIfFetchSuccessfull(args[0], ok, func() {
				endless := (len(args) > 1 && args[1] == "e")
				StartScreen(server, screens, config, endless)
			})
		case "stop":
			testIfFetchSuccessfull(args[0], ok, func() {
				StopScreen(server, screens, config)
			})
		case "resume":
			testIfFetchSuccessfull(args[0], ok, func() {
				ResumeScreen(server, screens, config)
			})
		case "restart":
			testIfFetchSuccessfull(args[0], ok, func() {
				endless := (len(args) > 1 && args[1] == "e")
				RestartScreen(server, screens, config, endless)
			})
		case "backup":
			testIfFetchSuccessfull(args[0], ok, func() {
				BackupMenu(server, config)
			})
		case "config":
			util.EditConfWithEditor(config, args[0])
		case "autostart":
			if args[0] == "reset" {
				if err := ResetAutostart(); err != nil {
					util.LogError("Failed deleting autostart file:\n" + err.Error())
				} else {
					util.LogInfo("Reset autostart.")
				}
				util.Pause()
				return
			}
			var selected []Screen
			for _, e := range args {
				if s, ok := fetchServer(servers, e); ok {
					selected = append(selected, *s)
				}
			}
			if len(selected) == 0 {
				util.LogError("No servers selected")
				util.Pause()
				return
			}
			err := CreateAutostart(&selected)
			if err == nil {
				util.LogInfo("Successfully added servers to autostart")
			} else {
				util.LogError("An unexpected error occured while creating autostart file:\n" + err.Error())
			}
			util.Pause()
		}
	}
}
