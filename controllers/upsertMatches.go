package controllers

import "../models"
import (
	"../models/match"
	"time"
)

func UpsertMatches(matches []models.Match) {
	time.Sleep(time.Millisecond * 500)
	for _, m := range matches {
		mtc := &match.Match{
			Matchid:               m.Matchid,
			Matchstatus:           m.Status,
			Matchtime:             m.Matchtime,
			Betstatus:             m.Betstatus,
			Score:                 m.Score,
			Msgnumber:             m.Msgnr,
			Earlybetstatus:        m.Earlybetstatus,
			Yrcardsaway:           m.Yellowcardsaway,
			Yrcardshome:           m.Yellowcardshome,
			Redcardsaway:          m.Redcardsaway,
			Redcardshome:          m.Redcardshome,
			Yellowcardsaway:       m.Yellowcardsaway,
			Yellowcardshome:       m.Yellowcardshome,
			Cornersaway:           m.Cornersaway,
			Cornershome:           m.Cornershome,
			Matchtimeextended:     m.MatchtimeExtended,
			Setscores:             m.Setscores,
			Active:                m.Active,
			Tiebreak:              m.Tiebreak,
			Service:               m.Server,
			Remainingtime:         m.Remainingtime,
			RemainingTimeinPeriod: m.Remainingtimeinperiod,
			Suspendaway:           m.SuspendAway,
			Suspendhome:           m.SuspendHome,
			Clockstop:             m.ClockStop,
			Gamescore:             m.Gamescore,
		}
		match.Model.Where(match.Match{Matchid: mtc.Matchid}).Assign(mtc).FirstOrCreate(mtc)
		//if len()
		if len(m.Odds) > 0 {
			go UpsertOdds(m)
		}
		if len(m.Card) > 0 {
			go UpsertCards(m)
		}
		if len(m.ScoreField) > 0 {
			go UpsertScores(m)
		}
	}
}
