package ws

import (
	"time"

	ws "github.com/mugsoft/tools/ws"
)

var Opts = &ws.Opts{
	Address:  ":1111",
	Time_out: time.Second * 2,
}

type Reply struct {
	Active      int           `json:"active"`
	Matchid     *int          `json:"matchid,omitempty"`
	Betstatus   *string       `json:"betstatus,omitempty"`
	Matchstatus *string       `json:"matchstatus,omitempty"`
	Service     *int          `json:"service,omitempty"`
	Score       *Score        `json:"score,omitempty"`
	Cards       *Cards        `json:"cards,omitempty"`
	ShownCards  []*ShownCards `json:"showncards,omitempty"`
	Time        *Time         `json:"time,omitempty"`
	Corners     *Corners      `json:"corners,omitempty"`
	Odds        []*Odd        `json:"odds,omitempty"`
}
type Corners struct {
	Home *int `json:"home,omitempty"`
	Away *int `json:"away,omitempty"`
}
type Time struct {
	Matchtime             *string `json:"matchtime,omitempty"`
	Remainingtime         *string `json:"remainingtime,omitempty"`
	RemainingTimeinPeriod *string `json:"remainingtimeinperiod,omitempty"`
	MatchtimeExtended     *string `json:"matchtimeextended,omitempty"`
	Clockstop             *int    `json:"clockstop,omitempty"`
}
type Odd struct {
	OddsId       *int        `json:"oddsid,omitempty"`
	OddsType     *int        `json:"oddstype,omitempty"`
	Special      *float64    `json:"special,omitempty"`
	Active       int         `json:"active,omitempty"`
	Typename     *string     `json:"typename,omitempty"`
	Mostbalanced *int        `json:"mostbalanced,omitempty"`
	Odds         []*OddField `json:"odds,omitempty"`
}
type OddField struct {
	Outcomeid *int     `json:"outcomeid,omitempty"`
	Active    int      `json:"active,omitempty"`
	Outcome   *string  `json:"outcome,omitempty"`
	Odd       *float64 `json:"odd,omitempty"`
}
type Score struct {
	Matchscore *string `json:"matchscore,omitempty"`
	Gamescore  *string `json:"gamescore,omitempty"`
	Setscores  *string `json:"setscores,omitempty"`
}
type Cards struct {
	SuspendAway   *int `json:"suspendaway,omitempty"`
	SuspendHome   *int `json:"suspendhome,omitempty"`
	Redhome       *int `json:"redhome,omitempty"`
	Yellowhome    *int `json:"yellowhome,omitempty"`
	Redaway       *int `json:"redaway,omitempty"`
	Yellowaway    *int `json:"yellowaway,omitempty"`
	Yellowredaway *int `json:"yellowredaway,omitempty"`
	Yellowredhome *int `json:"yellowredhome,omitempty"`
}
type ShownCards struct {
	Canceled *string `xml:"canceled,attr" json:"canceled,omitempty"`
	Id       *int    `xml:"id,attr" json:"id,omitempty"`
	Player   *string `xml:"player,attr" json:"player,omitempty"`
	Team     *string `xml:"team,attr" json:"team,omitempty"`
	Time     *int    `xml:"time,attr" json:"time,omitempty"`
	Type     *string `xml:"type,attr" json:"type,omitempty"`
	Playerid *int    `xml:"playerid,attr" json:"playerid,omitempty"`
}
