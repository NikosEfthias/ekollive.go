package repl

type Reply struct {
	Active      *int     `json:"active,omitempty"filter:"active"`
	Matchid     *int     `json:"matchid,omitempty"filter:"matchid"`
	Betstatus   *string  `json:"betstatus,omitempty"filter:"betstatus"`
	Matchstatus *string  `json:"matchstatus,omitempty"filter:"matchstatus"`
	Service     *int     `json:"service,omitempty"filter:"service"`
	Tiebreak    *string  `json:"tiebreak,omitempty"filter:"tiebreak"`
	Score       *Score   `json:"score,omitempty"`
	Cards       *Cards   `json:"cards,omitempty"`
	Time        *Time    `json:"time,omitempty"`
	Corners     *Corners `json:"corners,omitempty"`
	Odds        []*Odd   `json:"odds,omitempty"`
}
type Corners struct {
	Home *int `json:"home,omitempty"filter:"home"`
	Away *int `json:"away,omitempty"filter:"away"`
}
type Time struct {
	Matchtime             *string `json:"matchtime,omitempty"filter:"matchtime"`
	Remainingtime         *string `json:"remainingtime,omitempty"filter:"remainingtime"`
	RemainingTimeinPeriod *string `json:"remainingtimeinperiod,omitempty"filter:"remainingtimeinperiod"`
	MatchtimeExtended     *string `json:"matchtimeextended,omitempty"filter:"matchtimeextended"`
	Clockstop             *int    `json:"clockstop,omitempty"filter:"clockstop"`
}
type Odd struct {
	OddsId       int         `json:"oddsid,omitempty"filter:"oddsid"`
	OddsType     int         `json:"oddstype,omitempty"filter:"oddstype"`
	Special      *string     `json:"special,omitempty"filter:"special"`
	Active       *int        `json:"active,omitempty"filter:"odds.active"`
	Typename     *string     `json:"typename,omitempty"filter:"typename"`
	Mostbalanced *int        `json:"mostbalanced,omitempty"filter:"mostbalanced"`
	Odds         []*OddField `json:"odds,omitempty"`
}
type OddField struct {
	Outcomeid *int     `json:"outcomeid,omitempty"filter:"outcomeid"`
	Active    *int     `json:"active,omitempty"filter:"odds.odds.active"`
	Outcome   *string  `json:"outcome,omitempty"filter:"outcome"`
	Odd       *float64 `json:"odd,omitempty"filter:"odd"`
}
type Score struct {
	Matchscore *string `json:"matchscore,omitempty"filter:"matchscore"`
	Gamescore  *string `json:"gamescore,omitempty"filter:"gamescore"`
	Setscores  *string `json:"setscores,omitempty"filter:"setscores"`
}
type Cards struct {
	SuspendAway   *int `json:"suspendaway,omitempty"filter:"suspendaway"`
	SuspendHome   *int `json:"suspendhome,omitempty"filter:"suspendhome"`
	Redhome       *int `json:"redhome,omitempty"filter:"redhome"`
	Yellowhome    *int `json:"yellowhome,omitempty"filter:"yellowhome"`
	Redaway       *int `json:"redaway,omitempty"filter:"redaway"`
	Yellowaway    *int `json:"yellowaway,omitempty"filter:"yellowaway"`
	Yellowredaway *int `json:"yellowredaway,omitempty"filter:"yellowredaway"`
	Yellowredhome *int `json:"yellowredhome,omitempty"filter:"yellowredhome"`
}
