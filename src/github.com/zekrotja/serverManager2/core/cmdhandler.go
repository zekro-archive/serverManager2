package core

import (
	"os"
	"bufio"
	"fmt"
	"github.com/zekrotja/serverManager2/util"
)

func pause() {
	bufio.NewReader(os.Stdin).ReadString('\n')
}

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
		"\n exit                     | Exit the program")
	pause()
}

func HandleCmd(cmd string, screens []Screen, servers []Screen, config util.Conf ) {
	switch cmd {
	case "help":
		printHelp()
	}
}