//scraping/scraping
package main

import (
	"github.com/WinTicket/server/pkg/keirin/scraping"
	"github.com/WinTicket/server/pkg/keirin/scraping/kdreams"
	"github.com/WinTicket/server/pkg/keirin/scraping/keirinjp"
)

package scraping

import (
"context"
"github.com/WinTicket/server/pkg/keirin/scraping/kdreams"
"github.com/WinTicket/server/pkg/keirin/scraping/keirinjp"
"net/http"
)

type Client interface {
	FetchVenueInfos(ctx context.Context, ids []int64) ([]*VenueInfo, error)
}

type Scraping struct {
	keirinjp keirinjp.Client
	kdreams  kdreams.Client
	db *database.Database
}

type Params struct {
	KeirinjpHost string
	KdreamsHost  string
	db *database.Database
}

func NewClient(client *http.Client, params *Params) *Scraping {
	keCli := keirinjp.NewClient(client, params.KeirinjpHost)
	kdCli := kdreams.NewClient(client, params.KdreamsHost)
	return &Scraping{keirinjp: keCli, kdreams: kdCli}
}






// scraping/venue.go

package scraping

import (
"context"
"github.com/WinTicket/server/pkg/keirin/scraping/keirinjp"
)

func NewVenueInfo(parser *keirinjp.VenueJSONParser, histories []*keirinjp.WinnerHistory) (*VenueInfo, error) {
	second, err := parser.SecondToFloat64()
	if err != nil {
		return nil, err
	}
	date, err := parser.FormatDate()
	bestRecord := newBestRecord(parser.Data.PlayerName, second, date)

	factors := make([]*factor, 0)
	for _, parserFactor := range parser.Data.FirstFactor {
		percentage, err := parserFactor.PercentageToFloat64()
		if err != nil {
			return nil, err
		}

		factor := newFactor(parserFactor.Name, percentage, 1)
		factors = append(factors, factor)
	}

	for _, parserFactor := range parser.Data.SecondFactor {
		percentage, err := parserFactor.PercentageToFloat64()
		if err != nil {
			return nil, err
		}

		factor := newFactor(parserFactor.Name, percentage, 2)
		factors = append(factors, factor)
	}

	winnerHistories := make([]*winnerHistory, 0)
	for _, history := range histories {
		year, err := history.GetYear()
		if err != nil {
			return nil, err
		}
		win := newWinnerHistory(history.PlayerID, year)
		winnerHistories = append(winnerHistories, win)
	}

	return &VenueInfo{
		Factors:         factors,
		BestRecord:      bestRecord,
		WinerHistoryies: winnerHistories,
	}, nil
}

func (s *Scraping) FetchVenueInfo(ctx context.Context, ids []int64) ([]*VenueInfo, error) {
	venueInfos := make([]*VenueInfo, 0)
	for _, id := range ids {
		jsonParser, err := s.keirinjp.FetchVenueInfo(ctx, id)
		if err != nil {
			return nil, err
		}

		winnerHistories, err := s.keirinjp.FetchWinnerHistories(ctx, id)
		if err != nil {
			return nil, err
		}

		venueInfo, err := NewVenueInfo(jsonParser, winnerHistories)
		if err != nil {
			return nil, err
		}

		venueInfos = append(venueInfos, venueInfo)
	}

	return venueInfos, nil
}





// API/venue.go
func (k *keirinService) FetchVenueInfo(ctx context.Context, req *keirin.FetchVenueInfoRequest) (*keirin.FetchVenueInfoRequest, error) {
	venues, err := k.db.Venue.GetEnabled(ctx)
	if err != nil {
		return nil, togRPCError(err)
	}

	client := http.DefaultClient
	params := &scraping.Params{
		KeirinjpHost: keirinjp.Host,
		KdreamsHost:  kdreams.Host,
	}

	scrapingClient := scraping.NewScraping(client, params)
	err = scrapingClient.FetchVenueInfo(ctx, venues)
	if err != nil {
		return nil, togRPCError(err)
	}

	venueMap:=venues[0].ToUpdateMap()
	err = k.db.Venue.

		res := &keirin.FetchVenueInfoResponse{}
	return res, togRPCError(err)
}
