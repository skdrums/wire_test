package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	//"net/http"
	"strings"
)

type PlayerResult struct {
	EntryNumber     string
	Name            string
	Margin          string
	FinalHalfRecord string
	Factor          string
	Order           string
	FirstTop        bool // S
	FinalLapTop     bool // B
}

func main() {
	fileInfos, _ := ioutil.ReadFile("rakuten.html")
	fmt.Println(string(fileInfos))
	stringReader := strings.NewReader(string(fileInfos))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	playerResults := make([]*PlayerResult, 0, 9)
	doc.Find("#JS_CONTENTS_RESULT").Each(func(i int, s *goquery.Selection) {
		s.Find("table.result_table").Each(func(i int, s *goquery.Selection) {
			s.Find("tr").Each(func(i int, s *goquery.Selection) {
				if i == 0 {
					// ヘッダー部分なのでスキップ
					return
				}
				playerResult := &PlayerResult{}
				s.Find("td").Each(func(i int, s *goquery.Selection) {
					value := strings.TrimSpace(s.Text())
					switch i {
					case 1:
						playerResult.Order = value
						fmt.Println(value)
					case 2:
						playerResult.EntryNumber = value
					case 3:
						playerResult.Name = value
					case 4:
						playerResult.Margin = value
					case 5:
						playerResult.FinalHalfRecord = value
					case 6:
						playerResult.Factor = value
					case 7:
						if value == "SB" {
							playerResult.FirstTop = true
							playerResult.FinalLapTop = true
						}
						if value == "S" {
							playerResult.FirstTop = true
						}
						if value == "B" {
							playerResult.FinalLapTop = true
						}
					}
				})
				playerResults = append(playerResults, playerResult)
			})
		})
	})
}
