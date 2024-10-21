package main

import (
	"bufio"
	"os"
)

func Prompt() (string, error) {
	input := bufio.NewReader(os.Stdin)
	str, err := input.ReadString('\n')

	if err != nil {
		return "", err
	}

	return str, nil

}
