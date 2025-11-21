package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func Credentials(u string, p string) (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	username, err := getUsername(u, reader)
	if err != nil {
		return "", "", err
	}

	password, err := getPassword(p)
	if err != nil {
		return "", "", err
	}

	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}

func getUsername(u string, reader *bufio.Reader) (string, error) {
	if len(u) == 0 {
		fmt.Print("Username: ")
		username, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		return username, nil
	} else {
		return u, nil
	}
}

func getPassword(p string) (string, error) {
	if len(p) == 0 {
		fmt.Print("Password: ")
		bytePassword, err := term.ReadPassword(syscall.Stdin)
		fmt.Print("\n")
		if err != nil {
			return "", err
		}
		return string(bytePassword), nil
	} else {
		return p, nil
	}
}
