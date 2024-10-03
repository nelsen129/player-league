package store

import (
	"encoding/json"
	"fmt"
	"io"
)

// League type. For finding and sorting players
type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func (l League) Less(i, j int) bool {
	return l[i].Wins < l[j].Wins
}

func (l League) Len() int {
	return len(l)
}

func (l League) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		return nil, fmt.Errorf("problem parsing league, %v", err)
	}
	return league, nil
}
