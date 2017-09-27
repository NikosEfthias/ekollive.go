package controllers

import (
	"../models"
	"../models/card"
	"sync"
)

//insert card
var cardsLock sync.Mutex

func UpsertCards(match models.Match) {
	if len(match.Card) < 1 {
		return
	}
	for _, c := range match.Card {
		card.Model.Where(&card.Card{
			Cardid: c.Id,
		}).Assign(&card.Card{
			Cardid:      c.Id,
			Canceled:    c.Canceled,
			Player:      c.Player,
			Cardteam:    c.Team,
			Cardtime:    c.Time,
			Cardtype:    c.Type,
			Playerid:    c.Playerid,
			Matchid:     match.Matchid,
			Matchtime:   match.Matchtime,
			Matchscore:  match.Score,
			Matchstatus: match.Status,
		}).FirstOrCreate(&card.Card{})
	}
}
