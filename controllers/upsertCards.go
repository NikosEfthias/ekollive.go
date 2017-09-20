package controllers

import (
	"fmt"
	"../models"
	"../models/card"
)

//insert card

func UpsertCards(match models.Match) {
	if len(match.Card) < 1 {
		return
	}
	for _, c := range match.Card {
		fmt.Println(c)
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
