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
)

func branch() ([]byte, error) {
	c := exec.Command("git", "branch")
	return c.CombinedOutput()
}

func main() {
	fmt.Printf("gitsh version %s\n", version)

	reader := bufio.NewReader(os.Stdin)

	for {
		br, err := branch()
		brs := string(br)
		if err != nil {
			fmt.Println(brs)
			return
		}

		fmt.Printf("%s [%s] ", prompt, strings.TrimSpace(brs))

		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error reading command: %s\n", err)
			continue
		}

		args := strings.Split(cmd, " ")

		for i := range args {
			args[i] = strings.TrimSpace(args[i])
		}

		switch strings.ToLower(args[0]) {
		case "exit":
			return
		case "ng":
			if err := run(args[1:]); err != nil {
				fmt.Println(err)
			}
		default:
			args = append([]string{"git"}, args...)
			if err := run(args); err != nil {
				fmt.Println(err)
			}
		}

	}
}

func run(args []string) error {
	c := exec.Command(args[0], args[1:]...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
