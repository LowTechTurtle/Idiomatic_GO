package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	TeamName string
	Players  []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l League) MatchResult(Team1Name string, Team1Score int, Team2Name string, Team2Score int) {
	if _, ok := l.Wins[Team1Name]; !ok {
		return
	}

	if _, ok := l.Wins[Team2Name]; !ok {
		return
	}

	if Team1Score > Team2Score {
		l.Wins[Team1Name] += 1
	} else if Team2Score > Team1Score {
		l.Wins[Team2Name] += 1
	}
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Wins))
	for _, v := range l.Teams {
		names = append(names, v.TeamName)
	}

	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] < l.Wins[names[j]]
	})

	return names
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	res := r.Ranking()
	for _, v := range res {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}
func main() {
	l := League{
		Wins: map[string]int{},
		Teams: []Team{
			Team{"Italy", []string{"Player1", "Player2", "Player3"}},
			Team{"India", []string{"Player1", "Player2", "Player3"}},
			Team{"France", []string{"Player1", "Player2", "Player3"}},
			Team{"Niggeria", []string{"Player1", "Player2", "Player3"}},
		},
	}

	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Niggeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Niggeria", 110)
	l.MatchResult("Italy", 65, "Niggeria", 70)
	l.MatchResult("France", 95, "India", 80)
	RankPrinter(l, os.Stdout)
}
