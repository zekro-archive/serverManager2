package util

import (
	"fmt"
	"strings"
	"os"
	. "github.com/logrusorgru/aurora"
)


func LogInfo(msg string) {
	msgsplit := strings.Split(msg, "\n")
	for _, e := range msgsplit {
		fmt.Printf("%s | %s\n", Cyan("INFO"), e)
	}
}

func LogError(msg string) {
	msgsplit := strings.Split(msg, "\n")
	for _, e := range msgsplit {
		fmt.Printf("%s | %s\n", Red("ERR "), e)
	}
}

func LogWarn(msg string) {
	msgsplit := strings.Split(msg, "\n")
	for _, e := range msgsplit {
		fmt.Printf("%s | %s\n", Brown("WARN"), e)
	}
}

func LogFatal(msg string) {
	msgsplit := strings.Split(msg, "\n")
	for _, e := range msgsplit {
		fmt.Printf("%s | %s\n", Red("FATAL"), e)
	}
	os.Exit(-1)
}