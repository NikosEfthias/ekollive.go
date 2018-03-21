package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"../lib"
	"../lib/db"
	"../models"
	"../models/cancelbet"
	"../models/clearbet"
	"../models/match"
	"../models/odd"
	"../models/sportsbook"
)

var l sync.Mutex

func UpsertMatches(matches []models.Match, limiter chan bool, betradar models.BetradarLiveOdds) {
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
		switch strings.ToLower(*betradar.Status) {
		case "ended":
			fallthrough
		case "abandoned":
			fallthrough
		case "interrupted":
			fallthrough
		case "walkover":
			fallthrough
		case "walkover1":
			fallthrough
		case "walkover2":
			fallthrough
		case "retired":
			fallthrough
		case "retired1":
			fallthrough
		case "retired2":
			if *lib.Testing {
				fmt.Printf("match ended matchid=%d", *m.Matchid)
			}
			db.DB.DB().Exec("UPDATE odds SET active=0 where matchId=?", *m.Matchid)
		case "meta":
			fmt.Println("coming from meta", *m.MatchInfo.DateOfMatch, *m.MatchInfo.Category.Value)
			db.Upsert(db.DB2.DB(), sportsBook.Sport{}.Tablename(), &sportsBook.Sport{
				SportId:   m.MatchInfo.Sport.Id,
				SportName: m.MatchInfo.Sport.Value,
				ListOrder: m.MatchInfo.Sport.Id,
				Lang:      "en"})
			db.Upsert(db.DB2.DB(), sportsBook.Category{}.Tablename(), &sportsBook.Category{
				SportId:      m.MatchInfo.Sport.Id,
				Categoryid:   m.MatchInfo.Category.Id,
				CategoryName: m.MatchInfo.Category.Value,
				ListOrder:    m.MatchInfo.Category.Id,
				Lang:         "en",
			})
			db.Upsert(db.DB2.DB(), sportsBook.Tournament{}.Tablename(), &sportsBook.Tournament{
				SportId:        m.MatchInfo.Sport.Id,
				Categoryid:     m.MatchInfo.Category.Id,
				TournamentId:   m.MatchInfo.Tournament.Id,
				TournamentName: m.MatchInfo.Tournament.Value,
				ListOrder:      m.MatchInfo.Tournament.Id,
				Lang:           "en",
			})
			db.Upsert(db.DB2.DB(), sportsBook.Competitor{}.Tablename(), &sportsBook.Competitor{
				Lang:         "en",
				CompId:       m.MatchInfo.AwayTeam.Id,
				Comp2Id:      m.MatchInfo.AwayTeam.Uniqueid,
				SportId:      m.MatchInfo.Sport.Id,
				Categoryid:   m.MatchInfo.Category.Id,
				TournamentId: m.MatchInfo.Tournament.Id,
				CompName:     lib.Capitalize(m.MatchInfo.AwayTeam.Value),
			})
			db.Upsert(db.DB2.DB(), sportsBook.Competitor{}.Tablename(), &sportsBook.Competitor{
				Lang:         "en",
				CompId:       m.MatchInfo.HomeTeam.Id,
				Comp2Id:      m.MatchInfo.HomeTeam.Uniqueid,
				SportId:      m.MatchInfo.Sport.Id,
				Categoryid:   m.MatchInfo.Category.Id,
				TournamentId: m.MatchInfo.Tournament.Id,
				CompName:     lib.Capitalize(m.MatchInfo.HomeTeam.Value),
			})
			data := sportsBook.Match{
				SportId:      m.MatchInfo.Sport.Id,
				Categoryid:   m.MatchInfo.Category.Id,
				TournamentId: m.MatchInfo.Tournament.Id,
				Matchid:      m.Matchid,
				Comp1:        m.MatchInfo.HomeTeam.Id,
				Comp2:        m.MatchInfo.AwayTeam.Id,
				Matchdate:    time.Unix(int64(*m.MatchInfo.DateOfMatch)/1000, 0).Format("2006-01-02 15-04-05"),
				LiveActive:   "1",
				//PeriodLength:
			}
			mtc.SportId = m.MatchInfo.Sport.Id
			mtc.CategoryId = m.MatchInfo.Category.Id
			mtc.TournamentId = m.MatchInfo.Tournament.Id
			for _, extra := range m.MatchInfo.Infos {
				if extra.Type == "PeriodLength" && extra.Value != nil {

					i, err := strconv.Atoi(*extra.Value)
					if nil != err {
						fmt.Println("error", err, "extra info value:", *extra.Value)
						continue
					}
					data.PeriodLength = &i

				}
			}
			data.Update(db.DB2.DB())
		case "undocancelbet":
			oddIds := []*int{}
			for _, od := range m.Odds {
				oddIds = append(oddIds, od.Id)
			}
			cancelbet.Model.Where("oddid in ( ? )", oddIds).Delete(&odd.Odd{})
		case "cancelbet":
			oddIds := []*int{}
			for _, od := range m.Odds {
				val := &cancelbet.Cancelbet{
					Matchid:   m.Matchid,
					Oddid:     od.Id,
					Starttime: betradar.Starttime,
					Endtime:   betradar.Endtime,
				}
				db.Upsert(db.DB.DB(), "cancelbets", val)
				oddIds = append(oddIds, od.Id)
			}
			odd.Model.Where("oddid in ( ? )", oddIds).Update("active", 0)

		case "clearbet":
			for _, od := range m.Odds {
				for _, odf := range od.OddsField {
					db.Upsert(db.DB.DB(), "clearbets", &clearbet.Clearbet{
						Matchid:    m.Matchid,
						Oddid:      od.Id,
						Type:       odf.Type,
						Outcome:    odf.Outcome,
						Active:     odf.Active,
						VoidFactor: odf.Voidfactor,
					})
				}
			}
		default:
			if len(m.Odds) > 0 {
				UpsertOdds(m)
			}
			if len(m.Card) > 0 {
				UpsertCards(m)
			}
			if len(m.ScoreField) > 0 {
				UpsertScores(m)
			}
		}
		//upsert
		l.Lock()
		db.Upsert(db.DB.DB(), "matches", mtc)
		l.Unlock()
	}
	<-limiter
}
