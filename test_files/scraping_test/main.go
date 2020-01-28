package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
	//"os"
	//"strings"
)

type WinnerHistory struct {
	PlayerID string
	Date     string
}

func main() {
	//url := "https://keirin.jp/pc/dfw/portal/guest/data/winner_g3/winner_g3_hakodate.html"
	//resp, err := http.Get(url)
	//if err != nil {
	//	panic(err)
	//}
	//defer res.Body.Close()
	//doc, err := goquery.NewDocumentFromReader(res.Body)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	panic(err)
	//}

	v := &url.Values{}
	u := &url.URL{
		Scheme:   "http",
		Host:     "keirin.jp",
		Path:     "pc/dfw/portal/guest/data/winner_g3/winner_g3_hakodate.html",
		RawQuery: v.Encode(),
	}

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	ctx, _ := context.WithCancel(context.Background())
	client := http.DefaultClient
	res, err := client.Do(req.WithContext(ctx))
	if err != nil {
		fmt.Println("error1")
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println(res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("error2", err.Error())
	}
	fmt.Println("flag")

	winnerHistories := make([]*WinnerHistory, 0)
	table := doc.Find(`body > table:nth-child(8) > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(3) > td:nth-child(3) > table > tbody > tr > td > table`)
	table = table.Find("tr").Next().Each(func(i int, s *goquery.Selection) {
		var isNotCommemoration bool
		winnerHistory := &WinnerHistory{}
		s.Find("td").EachWithBreak(func(i int, s *goquery.Selection) bool {
			if isNotCommemoration {
				return true
			}
			value := s.Text()
			fmt.Println(value)
			switch i {
			case 0:
				if !strings.Contains(value, "周年") {
					isNotCommemoration = true
				}
			case 1:
				winnerHistory.Date = value
			case 2:
				winnerHistory.PlayerID = value
				winnerHistories = append(winnerHistories, winnerHistory)
				return false
			}
			return true
		})
	})
	//.Each(func(i int, s *goquery.Selection) {
	//fmt.Println(s.Text())
	fmt.Println(table.Html())
	//table.Find("tr").Each(func(i int, s *goquery.Selection) {
	//	if i == 0 {
	//		// ヘッダー部分なのでスキップ
	//		//fmt.Println(s.Text())
	//		return
	//	}
	//	//fmt.Println(s.Text())
	//	var isNotCommemoration bool
	//	winnerHistory := &WinnerHistory{}
	//	s.Find("td").Each(func(i int, s *goquery.Selection) {
	//		if isNotCommemoration {
	//			return
	//		}
	//		value := s.Text()
	//		switch i {
	//		case 0:
	//			if !strings.Contains(value, "周年") {
	//				isNotCommemoration = true
	//			}
	//		case 1:
	//			winnerHistory.Date = value
	//		case 2:
	//			winnerHistory.PlayerID = value
	//			winnerHistories = append(winnerHistories, winnerHistory)
	//		}
	//	})
	//})

	if len(winnerHistories) == 0 {
		fmt.Println("err")
	}

	for _, win := range winnerHistories {
		fmt.Println(fmt.Sprintf("%+v", *win))
	}
}
