package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	//"strconv"
)

//var (
//	venue = int64(11)
//)

type VenueJSONParser struct {
	Data     Data `json:"data"`
	ResultCd int  `json:"resultCd"`
}

//
//type Dataa struct {
//	Jyo string `json:"jyoCode"`
//}

type Data struct {
	PlayerName   string    `json:"maxAgariMei"`
	Second       string    `json:"maxAgari"`
	Date         string    `json:"kirokuDate"`
	FirstFactor  []*Factor `json:"firstTechniqList"`
	SecondFactor []*Factor `json:"secondTechniqList"`
}

type Factor struct {
	Name       string `json:"iconName"`
	Percentage string `json:"percentCnt"`
}

type RaceResult struct {
	Weather            string          `json:"tenki"`
	WindSpeedSrc       string          `json:"husoku"`
	PlayerResultsValid bool            `json:"tyakujyunDispFlg"`
	PlayerResults      []*PlayerResult `json:"tyakujyunItemSubData"`
	WonOdds            *WonOdds        `json:"haraiGakuSubData"`
}

type PlayerResult struct {
	OrderSrc           string `json:"tyaku"`
	EntryNumber        string `json:"syaban"`
	ID                 string `json:"sensyuRegistNo"`
	Name               string `json:"sensyuName"`
	MarginSrc          string `json:"tyakusa"`
	FactorSrc          string `json:"kimarite"`
	FinalHalfRecordSrc string `json:"agari"`
}

type WonOdds struct {
	BracketExacta   []*WinningTicket `json:"WT2HaraiGakuDispItemSubData"`
	BracketQuinella []*WinningTicket `json:"WH2HaraiGakuDispItemSubData"`
	Exacta          []*WinningTicket `json:"ST2HaraiGakuDispItemSubData"`
	Quinella        []*WinningTicket `json:"SH2HaraiGakuDispItemSubData"`
	QuinellaPlace   []*WinningTicket `json:"WHaraiGakuDispItemSubData"`
	Trifecta        []*WinningTicket `json:"RT3HaraiGakuDispItemSubData"`
	Trio            []*WinningTicket `json:"RH3HaraiGakuDispItemSubData"`
}

type WinningTicket struct {
	KeySrc    string `json:"kumiBan"`
	PayoffSrc string `json:"haraiGaku"`
}

func main() {
	//	v := &url.Values{}
	//	v.Set("jocd", venue)
	//v.Set("type", "JSJ020")
	//u := &url.URL{
	//	Scheme:   "http",
	//	Host:     "keirin.jp",
	//	Path:     "pc/json",
	//	RawQuery: v.Encode(),
	//}
	//
	//req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	//ctx, _ := context.WithCancel(context.Background())
	//client := http.DefaultClient
	//res, err := client.Do(req.WithContext(ctx))
	//res, err := http.Get("http://keirin.jp/pc/json?jocd=11&type=JSJ020")
	//if err != nil {
	//	return
	//}
	//defer res.Body.Close()
	//if res.StatusCode != http.StatusOK {
	//	fmt.Println("error")
	//	return
	//}
	//venueInfo := &VenueJSONParser{}
	//if err := json.NewDecoder(res.Body).Decode(venueInfo); err != nil {
	//	fmt.Println("error")
	//	return
	//}
	//fmt.Println(res.Body)
	//fmt.Println(fmt.Sprintf("%+v", venueInfo))

	id := "100"
	v := &url.Values{}
	v.Set("jocd", id)
	v.Set("type", "JSJ020")
	u := &url.URL{
		Scheme:   "http",
		Host:     "keirin.jp",
		Path:     "pc/json",
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
		fmt.Println("error2")
	}
	raceResult := &VenueJSONParser{}
	if err := json.NewDecoder(res.Body).Decode(raceResult); err != nil {
		fmt.Println("error3")
		return
	}
	fmt.Println(fmt.Sprintf("%+v", raceResult))
}
