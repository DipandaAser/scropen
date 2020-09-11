package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readKeyWord() (keywords string, err error) {
	var valid = false
	for !valid {
		fmt.Print("\n Enter a word for search (enter exit to exit) :")
		keywords, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		keywords = strings.TrimSpace(keywords)
		switch keywords {
		case "exit":
			fmt.Print("\n Good bye")
			os.Exit(1)
		case "":
			fmt.Print("\n KeyWords can't be blank")
			continue
		default:
			valid = true
		}
	}
	return keywords, nil
}
