package main

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Payload struct {
	Hash        string `json:"hash"`
	InitPayload bool   `json:"initPayload"`
	Params      struct {
		APISOURCE   string `json:"API_SOURCE"`
		Categoryid  string `json:"categoryid"`
		Day         string `json:"day"`
		InitPayload string `json:"initPayload"`
		Sport       string `json:"sport"`
		Token       string `json:"token"`
		Type        string `json:"type"`
		UUID        string `json:"uuid"`
	} `json:"params"`
	Result bool `json:"result"`
	Rows   struct {
		Soccer struct {
			League struct {
				AllowBets           bool   `json:"AllowBets"`
				AllowFixedEachWay   int    `json:"AllowFixedEachWay"`
				CanCashout          bool   `json:"CanCashout"`
				Category            string `json:"Category"`
				Comment             string `json:"Comment"`
				Country             string `json:"Country"`
				CountrySafe         string `json:"CountrySafe"`
				Description         string `json:"Description"`
				EventID             int    `json:"EventID"`
				EventStatus         string `json:"EventStatus"`
				HasActiveSubmarkets bool   `json:"HasActiveSubmarkets"`
				IgmAvailable        bool   `json:"IgmAvailable"`
				IsComp              bool   `json:"IsComp"`
				IsMatch             bool   `json:"IsMatch"`
				IsSuspended         bool   `json:"IsSuspended"`
				MarketCount         int    `json:"MarketCount"`
				Matches             []struct {
					AdditionalMarkets []int       `json:"AdditionalMarkets"`
					AllowBets         bool        `json:"AllowBets"`
					AllowFixedEachWay int         `json:"AllowFixedEachWay"`
					BetRule           interface{} `json:"BetRule"`
					CanCashout        bool        `json:"CanCashout"`
					Category          string      `json:"Category"`
					Comment           string      `json:"Comment"`
					Competitors       []struct {
						AllowBets         bool   `json:"AllowBets"`
						Draw              string `json:"Draw"`
						DrawID            string `json:"DrawID"`
						HasDrawOdds       bool   `json:"HasDrawOdds"`
						HasLineOdds       bool   `json:"HasLineOdds"`
						HasMargin1Odds    bool   `json:"HasMargin1Odds"`
						HasMargin2Odds    bool   `json:"HasMargin2Odds"`
						HasUnderOverOdds  bool   `json:"HasUnderOverOdds"`
						HasWinOdds        bool   `json:"HasWinOdds"`
						IsSuspended       bool   `json:"IsSuspended"`
						LevelOT           int    `json:"LevelOT"`
						Line              int    `json:"Line"`
						LineDiv           string `json:"LineDiv"`
						LineEventID       int    `json:"LineEventID"`
						LineID            string `json:"LineID"`
						Margin1Div        string `json:"Margin1Div"`
						Margin1ID         string `json:"Margin1ID"`
						Margin2Div        string `json:"Margin2Div"`
						Margin2ID         string `json:"Margin2ID"`
						Name              string `json:"Name"`
						PointsID          string `json:"PointsID"`
						Position          int    `json:"Position"`
						ScoreNT           int    `json:"ScoreNT"`
						ScoreOT           int    `json:"ScoreOT"`
						UnderOver         string `json:"UnderOver"`
						UnderOverHandicap int    `json:"UnderOverHandicap"`
						UnderOverID       string `json:"UnderOverID"`
						UnderOverString   string `json:"UnderOverString"`
						Win               string `json:"Win"`
						WinID             string `json:"WinID"`
					} `json:"Competitors"`
					Country                string      `json:"Country"`
					CountrySafe            string      `json:"CountrySafe"`
					Description            string      `json:"Description"`
					Draw                   string      `json:"Draw"`
					EventID                int         `json:"EventID"`
					EventStatus            string      `json:"EventStatus"`
					HasActiveSubmarkets    bool        `json:"HasActiveSubmarkets"`
					HasDrawOdds            bool        `json:"HasDrawOdds"`
					HasLineOdds            bool        `json:"HasLineOdds"`
					HasMargin1Odds         bool        `json:"HasMargin1Odds"`
					HasMargin2Odds         bool        `json:"HasMargin2Odds"`
					HasOdds                bool        `json:"HasOdds"`
					HasTotalPointsOdds     bool        `json:"HasTotalPointsOdds"`
					HasUnderOverOdds       bool        `json:"HasUnderOverOdds"`
					HasWinOdds             bool        `json:"HasWinOdds"`
					IgmAvailable           bool        `json:"IgmAvailable"`
					IsComp                 bool        `json:"IsComp"`
					IsMatch                bool        `json:"IsMatch"`
					IsSuspended            bool        `json:"IsSuspended"`
					Location               string      `json:"Location"`
					Margin1High            int         `json:"Margin1High"`
					Margin1Low             int         `json:"Margin1Low"`
					Margin2                int         `json:"Margin2"`
					MarketCount            int         `json:"MarketCount"`
					MatchHeading           string      `json:"MatchHeading"`
					NetBet                 string      `json:"NetBet"`
					NumberPlacings         int         `json:"NumberPlacings"`
					OutcomeDateTime        string      `json:"OutcomeDateTime"`
					OutcomeDateTimeInt     int         `json:"OutcomeDateTimeInt"`
					OverDiv                string      `json:"OverDiv"`
					OverallSortOrder       int         `json:"OverallSortOrder"`
					PerformData            bool        `json:"PerformData"`
					PerformVideo           bool        `json:"PerformVideo"`
					PhoneBetting           bool        `json:"PhoneBetting"`
					PlaceDivider           int         `json:"PlaceDivider"`
					QuickCallEnabled       bool        `json:"QuickCallEnabled"`
					RaceComment            string      `json:"RaceComment"`
					Sport                  string      `json:"Sport"`
					SportCSS               string      `json:"SportCSS"`
					Status                 string      `json:"Status"`
					SuspendDateTime        string      `json:"SuspendDateTime"`
					SuspendDateTimeInt     int         `json:"SuspendDateTimeInt"`
					TVChannel              bool        `json:"TVChannel"`
					TeamA                  string      `json:"TeamA"`
					TeamB                  string      `json:"TeamB"`
					Teams                  string      `json:"Teams"`
					TotalAdditionalMarkets int         `json:"TotalAdditionalMarkets"`
					TotalMarketCount       int         `json:"TotalMarketCount"`
					TotalOddsColumns       int         `json:"TotalOddsColumns"`
					TotalOver              int         `json:"TotalOver"`
					TotalUnder             int         `json:"TotalUnder"`
					Type                   string      `json:"Type"`
					UnderDiv               string      `json:"UnderDiv"`
					URL                    string      `json:"Url"`
					BetradarVideo          bool        `json:"betradarVideo"`
					BoostAvailable         bool        `json:"boostAvailable"`
					BoostType              interface{} `json:"boostType"`
					PromoType              interface{} `json:"promoType"`
					SuspendedMarkets       struct {
						Line        bool `json:"Line"`
						Margin1     bool `json:"Margin1"`
						Margin2     bool `json:"Margin2"`
						TotalPoints bool `json:"TotalPoints"`
						UnderOver   bool `json:"UnderOver"`
					} `json:"suspendedMarkets"`
				} `json:"Matches"`
				NetBet             string      `json:"NetBet"`
				OutcomeDateTime    string      `json:"OutcomeDateTime"`
				OutcomeDateTimeInt int         `json:"OutcomeDateTimeInt"`
				OverallSortOrder   int         `json:"OverallSortOrder"`
				PhoneBetting       bool        `json:"PhoneBetting"`
				PlaceDivider       int         `json:"PlaceDivider"`
				QuickCallEnabled   bool        `json:"QuickCallEnabled"`
				RaceComment        string      `json:"RaceComment"`
				SportCSS           string      `json:"SportCSS"`
				SuspendDateTime    string      `json:"SuspendDateTime"`
				SuspendDateTimeInt int         `json:"SuspendDateTimeInt"`
				URL                string      `json:"Url"`
				BetradarVideo      bool        `json:"betradarVideo"`
				BoostAvailable     bool        `json:"boostAvailable"`
				BoostType          interface{} `json:"boostType"`
				PromoType          interface{} `json:"promoType"`
			} `json:"League"`
		} `json:"Soccer"`
	} `json:"rows"`
	Sport       string `json:"sport"`
	Sql         []int  `json:"sql"`
	Status      string `json:"status"`
	TotalRows   int    `json:"totalRows"`
	TotalSports int    `json:"totalSports"`
	Type        string `json:"type"`
	UUID        string `json:"uuid"`
}

type Soccer struct {
	Modules         int       `json:"modules"`
	Pay             []Payload `json:"payload"`
	ServerTime      int       `json:"serverTime"`
	ServerTimestamp int       `json:"serverTimestamp"`
	ServerTimezone  string    `json:"serverTimezone"`
}

func SoccerScraper(URL string) (Soccer, error) {
	var data Soccer
	doc, err := goquery.NewDocument(URL)
	if err != nil {
		log.Fatal(err)
	}

	var t string
	doc.Find("script").Each(func(index int, item *goquery.Selection) {
		scrpt := item
		txt := scrpt.Text()
		if len(txt) > 13 && txt[:14] == "Delegator.init" {
			t = txt
		}
	})
	t, err = trimJSON(t)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(t), &data)
	if err != nil {
		log.Fatalln(err)
	}
	return data, nil
}

func trimJSON(a string) (string, error) {
	a = strings.TrimPrefix(a, "Delegator.init(")
	a = a[:len(a)-2]
	a = strings.Replace(a, ",[]", "", -1) // Faulty field ",[]" ocures in json extracted from HTML
	x := `"Category":"`
	y := `"`
	i := strings.Index(a, x)
	if i == -1 {
		return "", errors.New("page cannot be parsed")
	}

	b := a[i+len(x):]
	j := strings.Index(b, y)
	c := b[:j]

	return strings.Replace(a, c, "League", 1), nil
}
