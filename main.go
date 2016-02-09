package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	version = "0.1.0"
	prompt  = "g$"
	shell   = "sh"
	git     = "git"
)

func branch() ([]byte, error) {
	c := exec.Command(git, "branch")
	return c.CombinedOutput()
}

func main() {
	fmt.Printf("gitsh version %s\n", version)

	reader := bufio.NewReader(os.Stdin)

	for {
		br, err := branch()
		if err != nil {
			fmt.Println(string(br))
			return
		}

		fmt.Printf("%s [%s] ", prompt, string(br))

		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error reading command: %s\n", err)
			continue
		}

		args := strings.Split(cmd, " ")

		for i := range args {
			args[i] = strings.TrimSpace(args[i])
		}

		var proc string

		switch strings.ToLower(args[0]) {
		case "exit":
			return
		case "ng":
			proc = shell
			args = args[1:]
		default:
			proc = git
		}

		c := exec.Command(proc, args...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			fmt.Printf("error running command: %s\n", err)
			continue
		}

		fmt.Println(c.ProcessState)
	}
}
