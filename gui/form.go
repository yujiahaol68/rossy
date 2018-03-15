package gui

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	inputReader         *bufio.Reader
	ErrUserInvalidInput = errors.New("Invalid input")
)

func OptionalConfirm(message string, df bool) (bool, error) {
	switch df {
	case true:
		fmt.Printf("%s [y/n](y): ", message)
	case false:
		fmt.Printf("%s [y/n](n): ", message)
	}

	inputReader = bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		return df, err
	}

	switch strings.ToUpper(input) {
	case "Y":
		return true, nil
	case "N":
		return true, nil
	case "":
		return df, nil
	default:
		return df, ErrUserInvalidInput
	}
}

func OptionalInput(message, df string) (string, error) {
	inputReader = bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if input != "" {
		return df, nil
	}
	return input, nil
}
