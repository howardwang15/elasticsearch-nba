package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Attributes struct {
	Id                     int     `json:"id"`
	Name                   string  `json:"name"`
	Pos                    string  `json:"position"`
	Age                    int     `json:"age"`
	Team                   string  `json:"team"`
	GamesPlayed            int     `json:"games_played"`
	GamesStarted           int     `json:"games_started"`
	Minutes                float32 `json:"minutes"`
	FieldGoalsMade         float32 `json:"field_goals_made"`
	FieldGoalsAttempted    float32 `json:"field_goals_attempted"`
	FieldGoalsPercent      float32 `json:"field_goals_percentage"`
	ThreePointersMade      float32 `json:"three_point_made"`
	ThreePointersAttempted float32 `json:"three_point_attempted"`
	ThreePointersPercent   float32 `json:"three_point_percentage"`
	TwoPointersMade        float32 `json:"two_point_made"`
	TwoPointersAttempted   float32 `json:"two_point_attempted"`
	TwoPointersPercent     float32 `json:"two_point_percentage"`
	EffectiveFieldGoal     float32 `json:"effective_field_goal"`
	FreeThrowsMade         float32 `json:"free_throw_made"`
	FreeThrowsAttempted    float32 `json:"free_throw_attempted"`
	FreeThrowsPercent      float32 `json:"free_throw_percentage"`
	OffensiveRebounds      float32 `json:"offensive_rebounds"`
	DefensiveRebounds      float32 `json:"defensive_rebounds"`
	TotalRebounds          float32 `json:"total_rebounds"`
	Assists                float32 `json:"assists"`
	Steals                 float32 `json:"steals"`
	Blocks                 float32 `json:"blocks"`
	Turnovers              float32 `json:"turnovers"`
	Fouls                  float32 `json:"fouls"`
	Points                 float32 `json:"points"`
}

func main() {
	c := colly.NewCollector()
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		f, _ := os.Create("data.json")
		defer f.Close()

		e.ForEach("tr", func(playerIndex int, tr *colly.HTMLElement) {
			// if playerIndex == 2 {
			attributes := Attributes{}
			attributes.Id = playerIndex
			tr.ForEach("td", func(index int, td *colly.HTMLElement) {
				SetAttribute(&attributes, index, td.Text)
			})
			jsonAttributes, _ := json.Marshal(attributes)
			fmt.Println(string(jsonAttributes))
			f.WriteString(string(jsonAttributes) + "\n")
			// }
		})
	})
	c.OnHTML("tr.full_table.rowSum", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})
	c.OnHTML("table", func(e *colly.HTMLElement) {
		rows := e.Attr("tr[class=full_table rowSum]")
		fmt.Println(rows)
	})
	c.Visit("https://www.basketball-reference.com/leagues/NBA_2021_per_game.html")
	fmt.Println(c)
}
