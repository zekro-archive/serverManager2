package core

import (
	"fmt"
	"os"
	. "strings"
	"github.com/zekroTJA/serverManager2/util"
)

type Args struct {
	screens *[]Screen
	servers *[]Screen
	config  *util.Conf
	args    []string
}

// PRIVATE FUNCTIONS

func (self *Args) exists(invokes ...string) bool {
	for _, e := range self.args {
		for _, inv := range invokes {
			if e == inv {
				return true
			}
		}
	}
	return false
}

func (self *Args) getValue(invokes ...string) (string, bool) {
	for i, e := range self.args {
		for _, inv := range invokes {
			if e == inv && i+1 < len(self.args) {
				return self.args[i+1], true
			}
		}
	}
	return "", false
}

// PUBLIC FUNCTIONS

func (self *Args) Init(servers *[]Screen, screens *[]Screen, config *util.Conf) bool {
	self.args = os.Args
	self.config = config
	self.screens = screens
	self.servers = servers
	return len(self.args) > 1
}

func (self *Args) Parse(version string) {
	if self.exists("--help", "-h", "-?") {
		fmt.Println(
			"\n -s   --start      | Start (multiple) servers by NAME of the servers:" +
			"\n                   | Usage: -s server1,server2,..." + 
			"\n      --loop       | Use in combination with 'start' to start servers in loop" +
			"\n -t   --stop       | Stop (multiple) servers by NAME of the servers:" + 
			"\n                   | Usage: -t server1,server2,..." + 
			"\n -v   --version    | Get current programs version\n")
		return
	}

	if self.exists("--test") {
		self.config = util.GetConf("./testconf.json")
	}

	if self.exists("-v", "--version") {
		fmt.Println("ServerManager2 v." + version)
		return
	}

	if v, ok := self.getValue("-s", "--start"); ok {
		toStart := Split(v, ",")
		inLoop := self.exists("--loop")
		for _, e := range toStart {
			ok := StartScreen(
				&Screen { Name: e }, 
				self.screens,
				self.config,
				inLoop)
			if ok {
				util.LogInfo(
					fmt.Sprintf("Started server '%s'", e))
			} else {
				util.LogError(
					fmt.Sprintf("Failed starting server '%s'", e))
			}
		}
		return
	}

	if v, ok := self.getValue("-t", "--stop"); ok {
		toStart := Split(v, ",")
		for _, e := range toStart {
			ok := StopScreen(
				&Screen { Name: e }, 
				self.screens,
				self.config) 
			if ok {
				util.LogInfo(
					fmt.Sprintf("Stopped server '%s'", e))
			} else {
				util.LogError(
					fmt.Sprintf("Failed stopping server '%s'", e))
			}
		}
		return
	}
}