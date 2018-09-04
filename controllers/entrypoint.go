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
	"github.com/mugsoft/tools/ws"
	"github.com/nikosEfthias/ekollive.go/lib"
	"github.com/nikosEfthias/ekollive.go/models"
	wso "github.com/nikosEfthias/ekollive.go/ws"
)

var (
	compCounter = 0
	selCounter  = 0
)
var db *sql.DB
var comp__buffer []*models.BetconstructData
var sel__buffer []*models.BetconstructData
var Limiter = make(chan bool, *lib.J)

func init() {
	// go func() {
	// 	ticker := time.NewTicker(time.Millisecond * 200)
	// 	for {
	// 		<-ticker.C
	// 	}
	// }()
	_db, err := sql.Open("mysql", *lib.DB+"?parseTime=true&timeout=5s&writeTimeout=2s")
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
	case "GetSports":
		handle__GetSports(d)
	case "GetRegions":
		handle__GetRegions(d)
	case "GetCompetitions":
		compCounter++
		if compCounter < 500 {
			comp__buffer = append(comp__buffer, d)
			return
		}
		handle__GetCompetitions(comp__buffer)
		compCounter = 0
		comp__buffer = comp__buffer[:0]
		time.Sleep(time.Millisecond * 500)
	case "GetMarketTypes":
		handle__GetMarketTypes(d)
	case "GetSelectionTypes":
		selCounter++
		if selCounter < 500 {
			sel__buffer = append(sel__buffer, d)
			return
		}
		handle__GetSelectionTypes(sel__buffer)
		selCounter = 0
		sel__buffer = sel__buffer[:0]
		handle__GetSelectionTypes(sel__buffer)
		time.Sleep(time.Millisecond * 500)
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
	go __postData(*lib.PhpPostADDR, d)
}
func handle__GetSelectionTypes(d []*models.BetconstructData) {
	go __postData(*lib.PhpPostADDR, d)
}
func handle__SubscribePreMatch(d *models.BetconstructData) {

}
func handle__SubscribeMatches(d *models.BetconstructData) {
	if *d.Command == "SubscribeMatches" {
		Limiter <- true
		go sub__handle__match_statt(d)
		return
	}
	if d.Type == nil && *d.Command != "MatchStat" {
		return
	} else if d.Type == nil {
		Limiter <- true
		go sub__handle__match_statt(d)
		return
	}
	switch *d.Type {
	case "MatchStat":
		fallthrough
	case "Match":
		fallthrough
	case "Stat":
		Limiter <- true
		go sub__handle__match_statt(d)
	case "Market":
		Limiter <- true
		go sub__handle__match_market(d)
	}
}
func handle__SportUpdate(d *models.BetconstructData) {
	go __postData(*lib.PhpPostADDR, d)
}
func handle__RegionUpdate(d *models.BetconstructData) {
	go __postData(*lib.PhpPostADDR, d)
}
func handle__CompetitionUpdate(d *models.BetconstructData) {
	go __postData(*lib.PhpPostADDR, d)
}
func handle__MarketTypeUpdate(d *models.BetconstructData) {
	go __postData(*lib.PhpPostADDR, d)
}
func handle__TeamUpdate(d *models.BetconstructData) {
	go __postData(*lib.PhpPostADDR, d)
}
func handle__SelectionTypeUpdate(d *models.BetconstructData) {
	go __postData(*lib.PhpPostADDR, d)
}

func handle__GetSports(d *models.BetconstructData) {
	go __postData(*lib.PhpPostADDR, d)
}
func handle__GetCompetitions(d []*models.BetconstructData) {
	go __postData(*lib.PhpPostADDR, d)
}

func handle__GetRegions(d *models.BetconstructData) {
	go __postData(*lib.PhpPostADDR, d)
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
	if !*lib.IsLive {
		st, err := db.Prepare(queries["OddUpdate"])
		if nil != err {
			lib.Log_error(err)
			os.Exit(2)
		}
		for _, o := range d.Objects {
			for _, s := range o.Selections {
				if "MatchResult" == o.MarketKind {
					if nil != o.MatchId || nil != o.MarketTypeId {
						handle__betresult(s, *o.MatchId, *o.MarketTypeId)
					}
					continue
				}
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
					lib.Float_or_nil(o.Handicap),
					lib.Int_or_nil(s.SelectionTypeId),
					lib.Float_or_nil(s.Price),
					!is__suspended && is__visible,
					s.Price,
					time.Now().UTC().Format("2006-01-02 15:04:05"),
				)
				if nil != err {
					lib.Log_error(err)
					continue
				}
			}
		}
		st.Close()
		return
	}
	st, err := db.Prepare(queries["live_OddUpdate"])
	if nil != err {
		lib.Log_error(err)
		os.Exit(1)
	}
	for _, o := range d.Objects {
		var odd = &wso.Odd{
			Active:   lib.Bool_to_int(o.IsVisible),
			OddsId:   o.Id,
			OddsType: o.MarketTypeId,
			Special:  o.Handicap,
			Odds:     []*wso.OddField{},
		}
		for _, s := range o.Selections {
			if "MatchResult" == o.MarketKind {
				if nil != o.MatchId || nil != o.MarketTypeId {
					handle__betresult(s, *o.MatchId, *o.MarketTypeId)
				}
				continue
			}
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
			odf := &wso.OddField{
				Active:    lib.Bool_to_int(s.IsVisible),
				Outcomeid: s.SelectionTypeId,
				Odd:       s.Price,
			}
			odd.Odds = append(odd.Odds, odf)
			_, err := st.Exec(
				lib.Int_or_nil(s.Id),
				lib.Int_or_nil(o.MatchId),
				lib.Int_or_nil(s.SelectionTypeId),
				lib.Int_or_nil(o.MarketTypeId),
				lib.Float_or_nil(o.Handicap),
				lib.Float_or_nil(s.Price),
				!is__suspended && is__visible,
				s.Price,
				time.Now().UTC().Format("2006-01-02 15:04:05"),
			)
			if nil != err {
				lib.Log_error(err)
				continue
			}
		}
		if "MatchResult" == o.MarketKind {
			continue
		}
		dt, err := json.Marshal(wso.Reply{
			Active:  lib.Bool_to_int(o.IsVisible),
			Matchid: o.MatchId,
			Odds:    []*wso.Odd{odd},
		})
		if nil == err {
			ws.BroadcastJSON(&ws.Socket_data{Event: "data", Data: string(dt)}, wso.Opts)
			_ = dt
		}

	}
	st.Close()
	if nil != err {
		fmt.Fprintln(os.Stderr, err)
		lib.Log_error("error parsing competitors or something else", err)
	}
	return

}
func sub__handle__match_statt(d *models.BetconstructData) {
	defer done()
	if !*lib.IsLive {
		handle__pre_match_update(d)
		return
	}
	handle__live_match_update(d)
	return
}
func done() {
	<-Limiter
}
