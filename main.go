package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func isOpenSource(source string)  bool {
	openWords := []string{"open source","open-source","open code","opensource"}
	for _, words := range openWords{
		if !strings.Contains(source, words){
			continue
		}
		return true
	}
	return false
}

func findInGitHub(motcles string,maxPage int) []string {
	var liens []string
	if maxPage < 1 {
		maxPage=1
	}
	motcles = strings.ReplaceAll(motcles," ", "+")

	for currentPage := 1; currentPage <= maxPage; currentPage++ {
		sourceLink := fmt.Sprintf("https://github.com/search?p=%d&q=%s+open+source&type=Repositories", currentPage, motcles)

		res, err := http.Get(sourceLink)
		if err != nil {
			fmt.Println(err)
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Printf("Error when loading page : %v \n", sourceLink)
			return liens
		}


		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			log.Println("Error when parsing file")
		}

		fmt.Printf("\n Finding project on page : %v \n", sourceLink)

		doc.Find(".repo-list-item").Each(func(i int, selection *goquery.Selection) {
			lien, _ := selection.Find("a").Attr("href")
			if lien != "" {
				lien = "https://github.com" + lien


				// we go to the url and found if README.md Content the words open-source or similar words
				res, err := http.Get(lien)
				if err != nil {
					// error when loading page
					log.Print("")
				}

				if res.StatusCode != 200 {
					log.Printf("Error when loading page : %v \n ", lien)

				}

				doc, err := goquery.NewDocumentFromReader(res.Body)

				if err != nil {
					log.Println("Error when parsing file")
				}

				readmeContent := doc.Find(".Box-body").Text()

				if isOpenSource(readmeContent){
					fmt.Printf("\t Open source project found. Link : %v \n",lien)
					liens = append(liens, lien)
				}

			}

		})
	}
	return liens
}

func main() {
	words := "java"
	nbPages := 10

	fmt.Printf("Finding Open source Project in Github with keywords : %v \n", words)

	listeLiens := findInGitHub(words, 10)

	fmt.Printf("\n %d Open(s) source(s) Project(s) found in %d search pages Github with keywords : %v \n", len(listeLiens), nbPages, words)


}
