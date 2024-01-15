package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	FILE_NAME = "leaderboard.json"
)

type Score struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ByScore []Score

func (p ByScore) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p ByScore) Len() int           { return len(p) }
func (p ByScore) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type LeaderBoard struct {
	Leaders [100]Score
	Number  int
}

func (lb *LeaderBoard) AddToLeaderBoard(s Score) {
	lb.Leaders[lb.Number] = s
	lb.Number += 1
	lb.SortLeaders()
	lb.WriteLeaderBoardFile()
}

func (lb *LeaderBoard) SortLeaders() {
	sort.Sort(ByScore(lb.Leaders[0:lb.Number]))
}

func (lb *LeaderBoard) ReadLeaderBoardFile() {
	fileData, err := os.ReadFile(FILE_NAME)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	err = json.Unmarshal(fileData, &(lb.Leaders))
	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
		return
	}
	for i, leader := range lb.Leaders {
		if len(leader.Name) > 0 {
			lb.Number = i
		}
	}
}

func (lb *LeaderBoard) WriteLeaderBoardFile() {
	jsonData, err := json.Marshal(lb.Leaders[:lb.Number])
	if err != nil {
		fmt.Println("Error Marshaling JSON: ", err)
		return
	}
	err = os.WriteFile(FILE_NAME, jsonData, 0644)
}

func (lb *LeaderBoard) DrawLeaderBoard(s *ebiten.Image) {
	ebitenutil.DebugPrint(s, "Top Scores!!!")
	for i := 0; i < lb.Number; i++ {
		ebitenutil.DebugPrintAt(s, strconv.Itoa(i+1)+". "+lb.Leaders[lb.Number-i-1].Name+" "+strconv.Itoa(lb.Leaders[lb.Number-i-1].Value), 0, 11*(i+1))
	}
	ebitenutil.DebugPrintAt(s, "Press ESCAPE to restart game", 0, 11*(lb.Number+1))
}
