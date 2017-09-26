package websocketops

import (
	"github.com/gorilla/websocket"
	"sync"
	"../../models"
	"../../models/repl"
	"../store/oddids"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"../../lib"
	"strconv"
	"../store/filters"
)

var socketList [] *websocket.Conn
var l sync.Mutex

func Broadcast(data []byte) {
	l.Lock()
	defer l.Unlock()
	for _, c := range socketList {
		if nil != c.WriteMessage(websocket.TextMessage, data) {
			DelConnection(c)
		}
	}
}

func AddConnection(c *websocket.Conn) {
	l.Lock()
	defer l.Unlock()
	socketList = append(socketList, c)
}
func DelConnection(c *websocket.Conn) {
	l.Lock()
	defer l.Unlock()
	for i, sock := range socketList {
		if sock == c {
			socketList = append(socketList[:i], socketList[i+1:]...)
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
func StartBroadcast(c chan models.BetradarLiveOdds) {
	for d := range c {
		if *lib.Testing {
			goto testing
		}
		if len(socketList) == 0 {
			time.Sleep(time.Second)
			continue
		}
	testing:
		if checkStatuses(d) {
			//check match.Status to publish or not
			continue
		}
		for _, m := range d.Match {
			resp := &repl.Reply{
				Active:      m.Active,
				Matchid:     m.Matchid,
				Betstatus:   m.Betstatus,
				Matchstatus: m.Status,
				Service:     m.Server,
				Tiebreak:    m.Tiebreak,
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
				if odd.Type == nil && odd.Subtype == nil && odd.Typeid == nil {
					continue
				}
				o := &repl.Odd{
					OddsId:   *odd.Id,
					OddsType: oddids.Get(odd.Type, odd.Subtype, odd.Typeid),
					Special:  odd.Specialoddsvalue,
					Active:   odd.Active,
					Typename: odd.Freetext,
					Odds:     make([]*repl.OddField, 0),
				}
				for _, odf := range odd.OddsField {
					od_f := &repl.OddField{
						Outcomeid: odf.Typeid,
						Active:    odf.Active,
						Outcome:   odf.Type,
						Odd:       odf.InnerValue,
					}
					o.Odds = append(o.Odds, od_f)
				}
				resp.Odds = append(resp.Odds, o)
			}
			if resp.Matchid != nil {
				filters.ApplyFilters(resp, filters.GetFiltersByMatchId(strconv.Itoa(*resp.Matchid)))
			}
			dt, err := json.Marshal(resp)
			if nil != err {
				fmt.Println(err)
				continue
			}
			go Broadcast(dt)
		}
	}
}
