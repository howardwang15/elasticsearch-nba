package main

import "strconv"

func SetAttribute(attributes *Attributes, index int, value string) {
	switch index {
	case 0:
		attributes.Name = value
	case 1:
		attributes.Pos = value
	case 2:
		value, _ := strconv.Atoi(value)
		attributes.Age = value
	case 3:
		attributes.Team = value
	case 4:
		value, _ := strconv.Atoi(value)
		attributes.GamesPlayed = value
	case 5:
		value, _ := strconv.Atoi(value)
		attributes.GamesStarted = value
	case 6:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.Minutes = float32(value)
	case 7:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.FieldGoalsMade = float32(value)
	case 8:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.FieldGoalsAttempted = float32(value)
	case 9:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.FieldGoalsPercent = float32(value)
	case 10:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.ThreePointersMade = float32(value)
	case 11:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.ThreePointersAttempted = float32(value)
	case 12:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.ThreePointersPercent = float32(value)
	case 13:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.TwoPointersMade = float32(value)
	case 14:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.TwoPointersAttempted = float32(value)
	case 15:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.TwoPointersPercent = float32(value)
	case 16:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.EffectiveFieldGoal = float32(value)
	case 17:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.FreeThrowsMade = float32(value)
	case 18:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.FreeThrowsAttempted = float32(value)
	case 19:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.FreeThrowsPercent = float32(value)
	case 20:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.OffensiveRebounds = float32(value)
	case 21:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.DefensiveRebounds = float32(value)
	case 22:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.TotalRebounds = float32(value)
	case 23:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.Assists = float32(value)
	case 24:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.Steals = float32(value)
	case 25:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.Blocks = float32(value)
	case 26:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.Turnovers = float32(value)
	case 27:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.Fouls = float32(value)
	case 28:
		value, _ := strconv.ParseFloat(value, 32)
		attributes.Points = float32(value)
	}
}
