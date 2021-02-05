package main

type Player struct {
	Name     string  `json:"name"`
	GP       float32 `json:"GP,number"`
	MIN      float32 `json:"MIN"`
	PTS      float32 `json:"PTS"`
	FGM      float32 `json:"FGM"`
	FGA      float32 `json:"FGA"`
	FGPct    float32 `json:"FG%"`
	ThreePM  float32 `json:"3PM"`
	ThreePA  float32 `json:"3PA"`
	ThreePcg float32 `json:"3P%"`
	FTM      float32 `json:"FTM"`
	FTA      float32 `json:"FTA"`
	FTPcg    float32 `json:"FT%"`
	REB      float32 `json:"REB"`
	AST      float32 `json:"AST"`
	STL      float32 `json:"STL"`
	BLK      float32 `json:"BLK"`
	DD       float32 `json:"DD2"`
	TD       float32 `json:"TD3"`
	PER      float32 `json:"PER"`
}
