package util

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
)

func Cls() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Cinpt(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return text[0:len(text)-1]
}