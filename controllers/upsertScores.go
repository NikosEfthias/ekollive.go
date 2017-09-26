package controllers

import "../models"
import (
	"../models/score"
	"sync"
)

func UpsertScores(m models.Match, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, s := range m.ScoreField {
		score.Model.Where(&score.Score{
			Scoreid: s.Id,
		}).Assign(&score.Score{
			Scoreid:     s.Id,
			Home:        s.Home,
			Away:        s.Away,
			Player:      s.Player,
			Scoringteam: s.Scoringteam,
			Scoretime:   s.Time,
			Scoretype:   s.Type,
			Playerid:    s.Playerid,
			Matchid:     m.Matchid,
			Matchtime:   m.Matchtime,
			Matchscore:  m.Score,
			Matchstatus: m.Status,
		}).FirstOrCreate(&score.Score{})
	}
}
