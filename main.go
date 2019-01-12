package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Game struct is created in purpose of demonstration of SoccerScraper function.
// More deteiled struct can be obtained from source.
type Game struct {
	MatchName   string
	StartTime   time.Time
	Team1Name   string
	Team2Name   string
	WinTeam1    string
	WinTeam2    string
	Draw        string
	IsSuspended bool
}

// SoccerScraper pars only soccer/league pages on web www.ladbrokes.com.au
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

func examle(data Soccer) ([]Game, error) {
	var game Game
	gameSlice := make([]Game, len(data.Pay[0].Rows.Soccer.League.Matches))

	a := data.Pay[0].Rows.Soccer.League.Category
	for i, b := range data.Pay[0].Rows.Soccer.League.Matches {
		game.MatchName = a + ": " + b.Description
		game.Team1Name = b.Competitors[0].Name
		game.Team2Name = b.Competitors[1].Name
		if b.IsSuspended {
			game.WinTeam1 = "SUS"
			game.WinTeam2 = "SUS"
			game.Draw = "SUS"
		} else {
			game.WinTeam1 = b.Competitors[0].Win
			game.WinTeam2 = b.Competitors[1].Win
			game.Draw = b.Draw
		}
		game.IsSuspended = b.IsSuspended
		t, err := time.Parse(time.RFC1123Z, b.OutcomeDateTime)
		if err != nil {
			return nil, err
		}
		game.StartTime = t
		gameSlice[i] = game
	}
	return gameSlice, nil
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

func main() {
	m, err := SoccerScraper("https://www.ladbrokes.com.au/sports/soccer/69616684-football-australia-australian-a-league/")
	if err != nil {
		panic(err)
	}

	matches, err := examle(m)
	if err != nil {
		panic(err)
	}

	for _, m := range matches {
		p := fmt.Println
		p(m.MatchName)
		p(m.Team1Name)
		p(m.Team2Name)
		p(m.WinTeam1)
		p(m.WinTeam2)
		p(m.Draw)
		p(m.StartTime)
	}
}
