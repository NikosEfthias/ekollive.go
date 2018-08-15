package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/mugsoft/ekollive.go/lib"
	"github.com/mugsoft/ekollive.go/models"
)

var queries = map[string]string{
	"SubscribeMatches": "INSERT INTO `match`(`sportId`, `categoryId`, `tournamentId`, `matchId`, `manuelMatchId`, `comp1`, `comp2`, `matchDate`, `periodLength`, `liveActive`, `status`, `cancel`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE periodLength=?, matchDate=?, status=?, cancel=?, comp1 = ?, comp2 = ?, updatedAt = ?",

	"CompetitorUpdate": "INSERT  INTO competitor ( `sportId`, `categoryId`, `tournamentId`, `nameId`, `compId`, `compId2`, `compName`) VALUES (?,?,?,?,?,?,?) on duplicate key update `compName`=?,updatedAt=?",

	"MatchStatUpdate":      "insert into matchstats ( `matchId`,`name`,`iscanceled`,`istimeout`,`eventtimeutc`,`eventtype`,`eventtypeid`,`side`,`currentminute`, `matchlength`,`score`,`cornerscore`,`yellowcardscore`,`redcardscore`,`shotontargetscore`,`shotofftargetscore`,`dangerousattackscore`, `acesscore`,`doublefaultscore`,`sportkind`,`periodscore`,`quarterscore`,`setscore`,`set1score`,`set2score`,`set3score`, `set4score`,`set5score`,`set6score`,`set7score`,`set8score`,`set9score`,`set10score`,`gamescore`,`server`,`info`, `remainingtime`, `period`, `periodlength`, `setcount`, `penaltyscore`, `freekickscore`, `extratimescore`, `set1yellowcardscore`, `set2yellowcardscore`, `set1cornerscore`, `set2cornerscore`, `set1redcardscore`, `set2redcardscore`, `additionalminutes`, `teamId`) values( ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?, ?,?,?,?,?,?,?,?,?,?,?,?,?)",
	"live-MatchStatUpdate": "INSERT INTO `matches` (`matchid`, `gamescore`, `matchtime`, `matchstatus`, `betstatus`, `score`, `service`, `remainingtime`,  `redcardsaway`, `redcardshome`, `yellowcardsaway`, `yellowcardshome`, `cornersaway`, `cornershome`, `matchtimeextended`, `setscores`, `active`, `status`, `sportid`, `categoryid`, `tournamentid`) VALUES ( ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
	"live-MatchUpdate":     "INSERT INTO `matches` (`matchid`, `gamescore`, `matchtime`, `matchstatus`, `betstatus`, `score`, `service`, `remainingtime`,  `redcardsaway`, `redcardshome`, `yellowcardsaway`, `yellowcardshome`, `cornersaway`, `cornershome`, `matchtimeextended`, `setscores`, `active`, `status`, `sportid`, `categoryid`, `tournamentid`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
	"OddUpdate":            "INSERT INTO `odds`(`matchId`, `oddsType`, `outCome`, `special`, `outComeId`, `odds`, `status`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE odds=?, updatedAt=?",

	"live_OddUpdate": "INSERT INTO `odds`(`oddid`, `matchid`, `oddFieldTypeId`, `oddTypeId`, `specialvalue`,  `odd`, `active`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE odd= ?, updatedAt= ?",
}

func handle__pre_match_update(d *models.BetconstructData) {
	var comp1, comp2 int
	var comp1NameId, comp2NameId int
	var compName1, compName2 string
	var b__is_suspended, b__is_visible, ok bool
	var matchid int
	_ = ok
	var periodlength int
	var stats models.Stat
	for _, obj := range d.Objects {
		stats = models.Stat{}
		b__is_suspended = false
		b__is_visible = false
		periodlength = 0
		comp1, comp2 = 0, 0
		comp1NameId, comp2NameId = 0, 0
		compName1, compName2 = "", ""
		_membs := obj.MatchMembers
		stats = obj.Stat
		if stats.PeriodLength != nil {
			periodlength = *stats.PeriodLength
		}
		_, err := db.Exec(queries["MatchStatUpdate"],
			lib.Int_or_nil(stats.EventId),
			lib.String_or_nil(stats.Name),
			lib.Bool_or_nil(stats.IsCancelled),
			lib.Bool_or_nil(stats.IsTimeout),
			lib.Time_or_nil(stats.EventTimeUtc),
			lib.String_or_nil(stats.EventType),
			lib.Int_or_nil(stats.EventTypeId),
			lib.Int_or_nil(stats.Side),
			lib.Int_or_nil(stats.CurrentMinute),
			lib.Int_or_nil(stats.MatchLength),
			lib.String_or_nil(stats.Score),
			lib.String_or_nil(stats.CornerScore),
			lib.String_or_nil(stats.YellowcardScore),
			lib.String_or_nil(stats.RedcardScore),
			lib.String_or_nil(stats.ShotOnTargetScore),
			lib.String_or_nil(stats.ShotOffTargetScore),
			lib.String_or_nil(stats.DangerousAttackScore),
			lib.String_or_nil(stats.AcesScore),
			lib.String_or_nil(stats.DoubleFaultScore),
			lib.Int_or_nil(stats.SportKind),
			lib.String_or_nil(stats.PeriodScore),
			lib.String_or_nil(stats.QuarterScore),
			lib.String_or_nil(stats.SetScore),
			lib.String_or_nil(stats.Set1Score),
			lib.String_or_nil(stats.Set2Score),
			lib.String_or_nil(stats.Set3Score),
			lib.String_or_nil(stats.Set4Score),
			lib.String_or_nil(stats.Set5Score),
			lib.String_or_nil(stats.Set6Score),
			lib.String_or_nil(stats.Set7Score),
			lib.String_or_nil(stats.Set8Score),
			lib.String_or_nil(stats.Set9Score),
			lib.String_or_nil(stats.Set10Score),
			lib.String_or_nil(stats.GameScore),
			lib.Int_or_nil(stats.Server),
			lib.String_or_nil(stats.Info),
			lib.String_or_nil(stats.RemainingTime),
			lib.Int_or_nil(stats.Period),
			lib.Int_or_nil(stats.PeriodLength),
			lib.Int_or_nil(stats.SetCount),
			lib.String_or_nil(stats.PenaltyScore),
			lib.String_or_nil(stats.FreeKickScore),
			lib.String_or_nil(stats.ExtraTimeScore),
			lib.String_or_nil(stats.Set1YellowScore),
			lib.String_or_nil(stats.Set2YellowScore),
			lib.String_or_nil(stats.Set1CornerScore),
			lib.String_or_nil(stats.Set2CornerScore),
			lib.String_or_nil(stats.Set1RedScore),
			lib.String_or_nil(stats.Set2RedScore),
			lib.Int_or_nil(stats.AdditionalMinutes),
			lib.Int_or_nil(stats.TeamId),
		)
		if nil != err {
			fmt.Fprintln(os.Stderr, err)
		}

		for _, m := range _membs {
			var i__id int
			var s__compname string
			var ui__nameId uint
			var b__isHome bool
			if m.TeamId == nil {
				goto cont
			}
			i__id = *m.TeamId
			if m.IsHome == nil {
				goto cont
			}
			b__isHome = *m.IsHome
			if nil == m.NameId {
				goto cont
			}
			ui__nameId = *m.NameId
			if nil == m.TeamName {
				goto cont
			}
			s__compname = *m.TeamName
			if b__isHome {
				comp1 = i__id
				comp1NameId = int(ui__nameId)
				compName1 = s__compname
			} else {
				comp2 = i__id
				comp2NameId = int(ui__nameId)
				compName2 = s__compname
			}
		}
		_, err = db.Exec(queries["CompetitorUpdate"],
			lib.Int_or_nil(obj.SportId),
			lib.Int_or_nil(obj.RegionId),
			lib.Int_or_nil(obj.CompetitionId),
			comp1NameId,
			comp1,
			comp1,
			compName1,
			compName1,
			time.Now().UTC().Format("2006-01-02 15:04:05"),
		)
		if nil != err {
			goto cont
		}
		if comp1 != comp2 {
			_, err = db.Exec(queries["CompetitorUpdate"],
				lib.Int_or_nil(obj.SportId),
				lib.Int_or_nil(obj.RegionId),
				lib.Int_or_nil(obj.CompetitionId),
				comp2NameId,
				comp2,
				comp2,
				compName2,
				compName2,
				time.Now().UTC().Format("2006-01-02 15:04:05"),
			)
			if nil != err {
				goto cont
			}

		}
		if obj.IsVisible != nil {
			b__is_visible = *obj.IsVisible
		}

		//{{{
		if obj.IsSuspended != nil {
			b__is_suspended = *obj.IsSuspended
		}
		if obj.Id != nil {
			matchid = *obj.Id
		} else if obj.Stat.EventId != nil {
			matchid = *obj.Stat.EventId
		} else {
			continue
		}
		//}}}
		//{{{
		_, err = db.Exec(queries["SubscribeMatches"],
			lib.Int_or_nil(obj.SportId),
			lib.Int_or_nil(obj.RegionId),
			lib.Int_or_nil(obj.CompetitionId),
			matchid,
			matchid,
			comp1,
			comp2,
			lib.Time_or_nil(obj.Date),
			periodlength,
			lib.Bool_or_nil(obj.IsLive),
			!b__is_suspended && b__is_visible,
			b__is_suspended,
			periodlength,
			lib.Time_or_nil(obj.Date),
			!b__is_suspended && b__is_visible,
			b__is_suspended,
			comp1,
			comp2,
			time.Now().UTC().Format("2006-01-02 15:04:05"),
		)
		if nil != err {
			goto cont
		}
		continue
	cont:
		lib.Log_error("error parsing competitors or something else", err)
		continue
	}
	//}}}

}
func handle__live_match_update(d *models.BetconstructData) {
	var b__is_suspended, b__is_visible, ok bool
	_ = ok
	var stats models.Stat
	for _, obj := range d.Objects {
		stats = models.Stat{}
		b__is_suspended = false
		b__is_visible = false
		stats = obj.Stat
		if obj.IsVisible != nil {
			b__is_visible = *obj.IsVisible
		}
		if obj.IsSuspended != nil {
			b__is_suspended = *obj.IsSuspended
		}
		var periodLength *int = stats.PeriodLength
		if stats.PeriodLength == nil {
			periodLength = d.Stat.PeriodLength
		}

		var query = queries["live-MatchStatUpdate"] + lib.Prepare_on_duplicate_key_updates(map[string]interface{}{
			"gamescore":         stats.GameScore,
			"matchtime":         stats.EventTimeUtc,
			"matchstatus":       lib.Calculate_live_match_status(obj.MatchStatus, periodLength),
			"betstatus":         obj.LiveStatus,
			"score":             stats.Score,
			"service":           stats.Server,
			"remainingtime":     stats.RemainingTime,
			"redcardsaway":      lib.Split_score_fields(stats.RedcardScore, lib.Side_away),
			"redcardshome":      lib.Split_score_fields(stats.RedcardScore, lib.Side_home),
			"cornersaway":       lib.Split_score_fields(stats.CornerScore, lib.Side_away),
			"cornershome":       lib.Split_score_fields(stats.CornerScore, lib.Side_home),
			"matchtimeextended": stats.IsTimeout,
			"setscores":         stats.SetScore,
			"active":            obj.IsVisible,
			"status":            b__is_visible && !b__is_suspended,
			"sportid":           obj.SportId,
			"categoryid":        obj.RegionId,
			"tournamentid":      obj.CompetitionId,
		})
		_, err := db.Exec(query,
			lib.Int_or_nil(stats.EventId),
			lib.String_or_nil(stats.GameScore),
			lib.Time_or_nil(stats.EventTimeUtc),
			lib.String_or_nil(lib.Calculate_live_match_status(obj.MatchStatus, periodLength)),
			lib.Int_or_nil(obj.LiveStatus),
			lib.String_or_nil(stats.Score),
			lib.Int_or_nil(stats.Server), //not sure
			lib.String_or_nil(stats.RemainingTime),
			lib.Split_score_fields(stats.RedcardScore, lib.Side_away),
			lib.Split_score_fields(stats.RedcardScore, lib.Side_home),
			lib.Split_score_fields(stats.YellowcardScore, lib.Side_away),
			lib.Split_score_fields(stats.YellowcardScore, lib.Side_home),
			lib.Split_score_fields(stats.CornerScore, lib.Side_away),
			lib.Split_score_fields(stats.CornerScore, lib.Side_home),
			lib.Bool_or_nil(stats.IsTimeout),
			lib.String_or_nil(stats.SetScore),
			lib.Bool_or_nil(obj.IsVisible),
			b__is_visible && !b__is_suspended,
			lib.Int_or_nil(obj.SportId),
			lib.Int_or_nil(obj.RegionId),
			lib.Int_or_nil(obj.CompetitionId),
		)

		if nil != err {
			fmt.Fprintln(os.Stderr, err)
		}

		continue
		lib.Log_error("error parsing competitors or something else", err)
		continue
	}

}
