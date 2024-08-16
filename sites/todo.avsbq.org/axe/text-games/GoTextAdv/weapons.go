package main

import (
	"math/rand"
	"time"
)


type Weapon struct {
	minAtt int
	maxAtt int
	Name   string
}

func (w *Weapon) Fire() int {
	return w.minAtt + rand.Intn(w.maxAtt-w.minAtt)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

