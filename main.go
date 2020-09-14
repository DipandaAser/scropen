package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	file, _ := ioutil.ReadFile("tile.txt")
	title := string(file)
	fmt.Println(title)

	searchword, err := readKeyWord()
	if err != nil {
		fmt.Print("\n An error occured when reading input. Restart the program.")
		os.Exit(1)
	}

	var nbPages int

	fmt.Print("`\n In how much pages do you want to make search ? (leave blank for one) : ")
	fmt.Scanf("%d", &nbPages)
	if nbPages < 1 {
		nbPages = 1
	}

	fmt.Printf(" \t Finding Open source Project in Github with keywords : %v \n", searchword)

	listeLiens := findInGitHub(searchword, nbPages)

	fmt.Printf("\n %d Open(s) source(s) Project(s) found in %d search pages Github with keywords : %v \n", len(listeLiens), nbPages, searchword)
}
