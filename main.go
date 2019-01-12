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

// Game struct is created in purpose of demonstration of linkScrape function.
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

var data Data

func linkScrape(URL string) ([]Game, error) {

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
		return nil, err
	}

	err = json.Unmarshal([]byte(t), &data)
	if err != nil {
		log.Fatalln(err)
	}

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
	matches, err := linkScrape("https://www.ladbrokes.com.au/sports/soccer/69616684-football-australia-australian-a-league/")
	if err != nil {
		panic(err)
	}
	for _, m := range matches {
		fmt.Println(m.MatchName)
		fmt.Println(m.Team1Name)
		fmt.Println(m.Team2Name)
		fmt.Println(m.WinTeam1)
		fmt.Println(m.WinTeam2)
		fmt.Println(m.Draw)
	}
}
