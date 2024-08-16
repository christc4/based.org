package main


var Items = map[int]*Item{
	1: {Name: "Key"},
	2: {Name: "Chest", ItemForUse: 1, Contains: []int{3}},
	3: {Name: "Medal"},
}

var Weaps = map[int]*Weapon{
	1: {Name: "Phaser", minAtt: 5, maxAtt: 15},
	2: {Name: "Klingon Disruptor", minAtt: 1, maxAtt: 15},
	3: {Name: "Romulan Disruptor", minAtt: 3, maxAtt: 12},
}

var Enemies = map[int]*Character{
	0: {Name: "Klingon", Health: 50, Alive: true, Weap: 2, Npc: true},
	1: {Name: "Romulan", Health: 55, Alive: true, Weap: 3, Npc: true},
}


var evts = map[string]*Event{
	"alienAttack":     {Type: "Combat", Chance: 20, Description: "An alien beams in front of you and shoots you with a ray gun.", Health: -50, Evt: "doctorTreatment"},
	"doctorTreatment": {Type: "Story", Chance: 10, Description: "The doctor rushes in and inject you with a health boost.", Health: +30, Evt: ""},
	"android":         {Type: "Story", Chance: 50, Description: "Data is in the turbo lift and says hi to you", Health: 0, Evt: ""},
	"relaxing":        {Type: "Story", Chance: 100, Description: "In the lounge you are so relaxed that your health improves.", Health: +10, Evt: ""},
}


var LocationMap = map[string]*Location{
	"Bridge":      { Description: "You are on the bridge of a spaceship sitting in the Captain's chair.", Transitions: []string{"ReadyRoom", "TurboLift"}, Events: []string{"alienAttack"}},
	"ReadyRoom":  { Description: "The Captain's ready room.", Transitions: []string{"Bridge"}, Events: []string{}, Items: []int{2}},
	"TurboLift":  { Description: "A Turbo Lift that takes you anywhere in the ship.", Transitions: []string{"Bridge", "Lounge", "Engineering"}, Events: []string{"android"}},
	"Engineering": { Description: "You are in engineering where you see the star drive", Transitions: []string{"TurboLift"}, Events: []string{"alienAttack"}, Items: []int{1}},
	"Lounge":      { Description: "You are in the lounge, you feel very relaxed", Transitions: []string{"TurboLift"}, Events: []string{"relaxing"}},
}

