package main

import (
	"math/rand"
	"time"
)


type Event struct {
	Type        string
	Chance      int
	Description string
	Health      int
	Evt         string
}

func (e *Event) ProcessEvent(player *Character) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	if e.Chance >= r1.Intn(100) {
		if e.Type == "Combat" {
			//Generate opponent

			opp := new(Character)
			*opp = *Enemies[1+rand.Intn(len(Enemies)-1)]
			opp.Npc = true
			opp.Speed = 1 + rand.Intn(100)
			Output("green", "A " + opp.Name + " jumps in front of you and attacks")

			players := Players{*player, *opp}
			runBattle(players)
			// Run combat
			Output("green", "Combat Event")
		} else {
			Output("green", "\t" + e.Description)
			if e.Evt != "" {
				e.Health = e.Health + evts[e.Evt].ProcessEvent(player)
			}
		}
		return e.Health
	}
	return 0
}

