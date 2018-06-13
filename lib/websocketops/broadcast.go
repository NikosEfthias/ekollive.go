package websocketops

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"../../clientManager"
	"../../models"
	"../../models/repl"
	"../store/filters"
	"../store/oddids"
	"github.com/gorilla/websocket"
)

var SocketList []*websocket.Conn
var sockData = make(chan []byte)

func init() {
	go Broadcast(sockData)

}
func BroadCastNow(data []byte) {
	sockData <- data
}
func Broadcast(d chan []byte) {
	for {
		data := <-d
		go clientManager.Broadcast(data)
		for _, c := range SocketList {
			if nil != c.WriteMessage(websocket.TextMessage, data) {
				c.Close()
				DelConnection(c)
				continue
			}
		}
	}
}

func AddConnection(c *websocket.Conn) {
	SocketList = append(SocketList, c)
}
func DelConnection(c *websocket.Conn) {
	for i, sock := range SocketList {
		if sock == c {
			SocketList = append(SocketList[:i], SocketList[i+1:]...)
			break
		}
	}
}
func checkStatuses(data models.BetradarLiveOdds) bool {
	var statuses = []string{
		"change",
		"alive",
		"betstart",
		"betstop",
	}
	for _, d := range statuses {
		if strings.ToLower(*data.Status) == d {
			return false
		}
	}
	return true
}
func StartBroadcast(c chan *models.BetradarLiveOdds) {
	for d := range c {
		if d == nil || checkStatuses(*d) {
			//check match.Status to publish or not
			continue
		}
		if *d.Status == "alive" && len(d.Match) == 0 {
			dt, err := json.Marshal(d)
			if nil != err {
				fmt.Println(err)
				continue
			}
			sockData <- dt
			continue
		}
		for _, m := range d.Match {
			resp := &repl.Reply{
				Active:         m.Active,
				Matchid:        m.Matchid,
				Betstatus:      m.Betstatus,
				Matchstatus:    m.Status,
				Earlybetstatus: m.Earlybetstatus,
				Service:        m.Server,
				Tiebreak:       m.Tiebreak,
				Score: &repl.Score{
					Matchscore: m.Score,
					Gamescore:  m.Gamescore,
					Setscores:  m.Setscores,
				},
				Cards: &repl.Cards{
					SuspendAway:   m.SuspendAway,
					SuspendHome:   m.SuspendHome,
					Redaway:       m.Redcardsaway,
					Redhome:       m.Redcardshome,
					Yellowaway:    m.Yellowcardsaway,
					Yellowhome:    m.Yellowcardshome,
					Yellowredaway: m.Yellowredcardsaway,
					Yellowredhome: m.Yellowredcardshome,
				},
				Time: &repl.Time{
					Matchtime:             m.Matchtime,
					Remainingtime:         m.Remainingtime,
					RemainingTimeinPeriod: m.Remainingtimeinperiod,
					MatchtimeExtended:     m.MatchtimeExtended,
					Clockstop:             m.ClockStop,
				},
				Corners: &repl.Corners{
					Home: m.Cornershome,
					Away: m.Cornersaway,
				},
				Odds: make([]*repl.Odd, 0),
			}
			for _, odd := range m.Odds {

				o := &repl.Odd{
					OddsId:       *odd.Id,
					OddsType:     oddids.Get(odd.Type, odd.Subtype, odd.Typeid),
					Special:      odd.Specialoddsvalue,
					Active:       odd.Active,
					Typename:     odd.Freetext,
					Mostbalanced: odd.Mostbalanced,
					Odds:         make([]*repl.OddField, 0),
				}
				if nil != odd.OddTypeId {
					o.OddsType = *odd.OddTypeId
				}
				for _, odf := range odd.OddsField {
					od_f := &repl.OddField{
						Outcomeid: odf.Typeid,
						Active:    odf.Active,
						Outcome:   odf.Type,
					}
					if odf.InnerValue != nil && *odf.InnerValue != "" {
						f, err := strconv.ParseFloat(*odf.InnerValue, 64)
						if nil != err {
							fmt.Println("cannot parse odd odd value=", f, err)
							return
						}
						od_f.Odd = &f
					}
					o.Odds = append(o.Odds, od_f)
				}
				resp.Odds = append(resp.Odds, o)
			}
			if resp.Matchid != nil {
				filters.ApplyFilters(resp, filters.GetFiltersByMatchId(strconv.Itoa(*resp.Matchid)))
			}
			dt, err := json.Marshal(resp)
			//dx, err := json.MarshalIndent(resp, "", "	")
			if nil != err {
				fmt.Println(err)
				continue
			}
			//fmt.Println(string(dx))
			sockData <- dt
		}
	}
}
