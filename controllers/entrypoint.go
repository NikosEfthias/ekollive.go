package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/pp"
	"github.com/mugsoft/ekollive.go/lib"
	"github.com/mugsoft/ekollive.go/models"
)

var (
	compCounter = 0
	selCounter  = 0
)
var db *sql.DB
var comp__buffer []*models.BetconstructData
var sel__buffer []*models.BetconstructData
var limiter = make(chan bool, 150)

func init() {
	go func() {
		ticker := time.NewTicker(time.Millisecond * 200)
		for {
			<-ticker.C
			fmt.Println("	||	", len(limiter), "	||	")
		}
	}()
	_db, err := sql.Open("mysql", "nikos:12345611@tcp(18.184.217.74:3306)/sportsdata?parseTime=true&timeout=5s&writeTimeout=2s")
	if nil != err {
		log.Fatalln(err)
	}
	err = _db.Ping()
	if nil != err {
		log.Println(err)
		log.Fatalln("baglanmiyor")
	}
	db = _db
	comp__buffer = make([]*models.BetconstructData, 0, 500)
	sel__buffer = make([]*models.BetconstructData, 0, 500)

}

var queries = map[string]string{
	"SubscribeMatches": "INSERT INTO `match`(`sportId`, `categoryId`, `tournamentId`, `matchId`, `manuelMatchId`, `comp1`, `comp2`, `matchDate`, `periodLength`, `liveActive`, `status`, `cancel`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE periodLength=?, matchDate=?, status=?, cancel=?, comp1 = ?, comp2 = ?, updatedAt = ?",

	"CompetitorUpdate": "INSERT  INTO competitor ( `sportId`, `categoryId`, `tournamentId`, `nameId`, `compId`, `compId2`, `compName`) VALUES (?,?,?,?,?,?,?) on duplicate key update `compName`=?,updatedAt=?",

	"MatchStatUpdate": "insert into matchstats ( `matchId`,`name`,`iscanceled`,`istimeout`,`eventtimeutc`,`eventtype`,`eventtypeid`,`side`,`currentminute`, `matchlength`,`score`,`cornerscore`,`yellowcardscore`,`redcardscore`,`shotontargetscore`,`shotofftargetscore`,`dangerousattackscore`, `acesscore`,`doublefaultscore`,`sportkind`,`periodscore`,`quarterscore`,`setscore`,`set1score`,`set2score`,`set3score`, `set4score`,`set5score`,`set6score`,`set7score`,`set8score`,`set9score`,`set10score`,`gamescore`,`server`,`info`, `remainingtime`, `period`, `periodlength`, `setcount`, `penaltyscore`, `freekickscore`, `extratimescore`, `set1yellowcardscore`, `set2yellowcardscore`, `set1cornerscore`, `set2cornerscore`, `set1redcardscore`, `set2redcardscore`, `additionalminutes`, `teamId`) values( ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?, ?,?,?,?,?,?,?,?,?,?,?,?,?)",
	"OddUpdate":       "INSERT INTO `odds`(`matchId`, `oddsType`, `outCome`, `special`, `outComeId`, `odds`, `status`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE odds=?, updatedAt=?",
}

func Handle_data(d *models.BetconstructData) {
	if *d.Command != "GetCompetitions" && compCounter > 0 {
		compCounter = 0
		handle__GetCompetitions(comp__buffer)
		comp__buffer = comp__buffer[:0]
	}
	if *d.Command != "GetSelectionTypes" && selCounter > 0 {
		selCounter = 0
		handle__GetSelectionTypes(sel__buffer)
		sel__buffer = sel__buffer[:0]
	}
	switch *d.Command {
	case "HeartBeat":
		pp.Println(d)
	case "GetSports":
		handle__GetSports(d)
	case "GetRegions":
		handle__GetRegions(d)
	case "GetCompetitions":
		compCounter++
		pp.Println(compCounter)
		if compCounter < 500 {
			comp__buffer = append(comp__buffer, d)
			pp.Println(compCounter)
			return
		}
		fmt.Println("waiting")
		handle__GetCompetitions(comp__buffer)
		compCounter = 0
		comp__buffer = comp__buffer[:0]
		time.Sleep(time.Millisecond * 500)
	case "GetMarketTypes":
		handle__GetMarketTypes(d)
	case "GetMatches":
		handle__GetMatches(d)
	case "GetSelectionTypes":
		selCounter++
		pp.Println(selCounter)
		if selCounter < 500 {
			sel__buffer = append(sel__buffer, d)
			return
		}
		fmt.Println("waiting")
		handle__GetSelectionTypes(sel__buffer)
		selCounter = 0
		sel__buffer = sel__buffer[:0]
		time.Sleep(time.Millisecond * 500)
		handle__GetSelectionTypes(sel__buffer)
	case "SubscribePreMatch":
		handle__SubscribePreMatch(d)
	case "MatchUpdate":
		fallthrough
	case "MatchStat":
		fallthrough
	case "SubscribeMatches":
		handle__SubscribeMatches(d)
	case "SportUpdate":
		handle__SportUpdate(d)
	case "RegionUpdate":
		handle__RegionUpdate(d)
	case "CompetitionUpdate":
		handle__CompetitionUpdate(d)
	case "MarketTypeUpdate":
		handle__MarketTypeUpdate(d)
	case "TeamUpdate":
		handle__TeamUpdate(d)
	case "SelectionTypeUpdate":
		handle__SelectionTypeUpdate(d)
	}
}

func handle__GetMarketTypes(d *models.BetconstructData) {
	go __postData("http://18.184.217.74/parseme", d)
}
func handle__GetMatches(d *models.BetconstructData) {
}
func handle__GetSelectionTypes(d []*models.BetconstructData) {
	go __postData("http://18.184.217.74/parseme", d)
}
func handle__SubscribePreMatch(d *models.BetconstructData) {

}
func handle__MatchUpdate(d *models.BetconstructData) {

}
func handle__MatchStat(d *models.BetconstructData) {

}
func handle__SubscribeMatches(d *models.BetconstructData) {
	limiter <- true

	if d.Type == nil && *d.Command != "MatchStat" {
		done()
		return
	} else if d.Type == nil {
		// pp.Println(d)
		done()
		return
	}
	switch *d.Type {
	case "Match":
		go sub__handle__match_statt(d)
	case "Stat":
		go sub__handle__match_statt(d)
	case "Market":
		go sub__handle__match_market(d)
	}
}
func handle__SportUpdate(d *models.BetconstructData) {
	go __postData("http://18.184.217.74/parseme", d)
}
func handle__RegionUpdate(d *models.BetconstructData) {
	go __postData("http://18.184.217.74/parseme", d)
}
func handle__CompetitionUpdate(d *models.BetconstructData) {
	go __postData("http://18.184.217.74/parseme", d)
}
func handle__MarketTypeUpdate(d *models.BetconstructData) {
	go __postData("http://18.184.217.74/parseme", d)
}
func handle__TeamUpdate(d *models.BetconstructData) {
	go __postData("http://18.184.217.74/parseme", d)
}
func handle__SelectionTypeUpdate(d *models.BetconstructData) {
	go __postData("http://18.184.217.74/parseme", d)
}

func handle__GetSports(d *models.BetconstructData) {
	go __postData("http://18.184.217.74/parseme", d)
}
func handle__GetCompetitions(d []*models.BetconstructData) {
	go __postData("http://18.184.217.74/parseme", d)
}

func handle__GetRegions(d *models.BetconstructData) {
	go __postData("http://18.184.217.74/parseme", d)
}
func __postData(addr string, data interface{}) {
	vals := url.Values{}
	dt, err := json.Marshal(data)
	if nil != err {
		return
	}
	vals.Set("data", string(dt))
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.PostForm(addr, vals)
	if nil != err {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Fprintln(os.Stderr, err)
	} else if resp.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, string(body), resp.StatusCode)
	}
}
func sub__handle__match_market(d *models.BetconstructData) {
	defer done()
	st, err := db.Prepare(queries["OddUpdate"])
	if nil != err {
		lib.Log_error(err)
	}
	for _, o := range d.Objects {
		for _, s := range o.Selections {
			if nil == s.SelectionTypeId {
				continue
			}
			var is__visible, is__suspended bool
			if s.IsSuspended != nil {
				is__suspended = *s.IsSuspended
			}
			if s.IsVisible != nil {
				is__visible = *s.IsVisible
			}
			_, err := st.Exec(lib.Int_or_nil(o.MatchId),
				lib.Int_or_nil(o.MarketTypeId),
				lib.String_or_nil(s.Name),
				lib.Int_or_nil(o.Handicap),
				lib.Int_or_nil(s.SelectionTypeId),
				lib.Float_or_nil(s.Price),
				!is__suspended && is__visible,
				s.Price,
				time.Now().UTC().Format("2006-01-02 15:04:05"),
			)
			if nil != err {
				pp.Println(d)
				lib.Log_error(err)
				continue
			}
		}
	}
	st.Close()
}
func sub__handle__match_statt(d *models.BetconstructData) {
	defer done()
	var comp1, comp2 int
	var comp1NameId, comp2NameId int
	var compName1, compName2 string
	var b__is_suspended, b__is_visible, ok bool
	var matchid int
	_ = ok
	var periodlength int
	var stats *models.Stat
	for _, obj := range d.Objects {
		stats = nil
		b__is_suspended = false
		b__is_visible = false
		periodlength = 0
		comp1, comp2 = 0, 0
		comp1NameId, comp2NameId = 0, 0
		compName1, compName2 = "", ""
		_membs := obj.MatchMembers
		stats = obj.Stat
		if stats != nil && stats.PeriodLength != nil {
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
			lib.Int_or_nil(stats.RemainingTime),
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
			fmt.Println(err)
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
			pp.Println(err)
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
				pp.Println(err)
				goto cont
			}

		}
		if obj.IsVisible != nil {
			b__is_visible = *obj.IsVisible
		}

		if obj.IsSuspended != nil {
			b__is_suspended = *obj.IsSuspended
		}
		if obj != nil && obj.Id != nil {
			matchid = *obj.Id
		} else if obj.Stat != nil && obj.Stat.EventId != nil {
			matchid = *obj.Stat.EventId
		} else {
			continue
		}

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
			pp.Println(err)
			os.Exit(1)
			goto cont
		}
		continue
	cont:
		lib.Log_error("error parsing competitors or something else", err)
		continue
	}
}
func done() {
	<-limiter
}
