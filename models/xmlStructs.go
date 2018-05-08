package models

type BetradarLiveOdds struct {
	Timestamp *int    `xml:"timestamp,attr"json:"timestamp,omitempty"`
	Status    *string `xml:"status,attr"json:"status,omitempty"`
	Time      *string `xml:"time,attr"json:"time,omitempty"`
	Replytype *string `xml:"replytype,attr"json:"replytype,omitempty"`
	Starttime *int    `xml:"starttime,attr"json:"starttime,omitempty"`
	Endtime   *int    `xml:"endtime,attr"json:"endtime,omitempty"`
	Match     []Match `xml:"Match"json:"match,omitempty"`
}

type Match struct {
	Remainingtimeinperiod *string   `xml:"remaining_time_in_periodd,attr"json:"remainingtimeinperiodd,omitempty"`
	Remainingtime         *string   `xml:"remaining_time,attr"json:"remainingtime,omitempty"`
	SuspendHome           *int      `xml:"suspendHome,attr"json:"suspendHome,omitempty"`
	SuspendAway           *int      `xml:"suspendAway,attr"json:"suspendAway,omitempty"`
	ClockStop             *int      `xml:"clock_stop,attr"json:"clock_stop,omitempty"`
	Matchid               *int      `xml:"matchid,attr"json:"matchid,omitempty"`
	Gamescore             *string   `xml:"gamescore,attr"json:"gamescore,omitempty"`
	Server                *int      `xml:"server,attr"json:"server,omitempty"`
	Tiebreak              *string   `xml:"tiebreak,attr"json:"tiebreak,omitempty"`
	Matchtime             *string   `xml:"matchtime,attr"json:"matchtime,omitempty"`
	Status                *string   `xml:"status,attr"json:"status,omitempty"`
	Betstatus             *string   `xml:"betstatus,attr"json:"betstatus,omitempty"`
	Score                 *string   `xml:"score,attr"json:"score,omitempty"`
	Msgnr                 *int      `xml:"msgnr,attr"json:"msgnr,omitempty"`
	Earlybetstatus        *string   `xml:"earlybetstatus,attr"json:"earlybetstatus,omitempty"`
	Active                *int      `xml:"active,attr"json:"active,omitempty"`
	Setscores             *string   `xml:"setscores,attr"json:"setscores,omitempty"`
	Cornersaway           *int      `xml:"cornersaway,attr"json:"cornersaway,omitempty"`
	Cornershome           *int      `xml:"cornershome,attr"json:"cornershome,omitempty"`
	Yellowredcardsaway    *int      `xml:"yellowredcardsaway,attr"json:"yellowredcardsaway,omitempty"`
	Yellowredcardshome    *int      `xml:"yellowredcardshome,attr"json:"yellowredcardshome,omitempty"`
	Redcardsaway          *int      `xml:"redcardsaway,attr"json:"redcardsaway,omitempty"`
	Redcardshome          *int      `xml:"redcardshome,attr"json:"redcardshome,omitempty"`
	Yellowcardsaway       *int      `xml:"yellowcardsaway,attr"json:"yellowcardsaway,omitempty"`
	Yellowcardshome       *int      `xml:"yellowcardshome,attr"json:"yellowcardshome,omitempty"`
	MatchtimeExtended     *string   `xml:"matchtime_extended,attr"json:"matchtime_extended,omitempty"`
	ScoreField            []Score   `xml:"Score"json:"score_field,omitempty"`
	Card                  []Card    `xml:"Card"json:"card,omitempty"`
	Odds                  []Odd     `xml:"Odds"json:"odds,omitempty"`
	MatchInfo             MatchInfo `xml:"MatchInfo"json:"MatchInfo,omitempty"`
}
type MatchInfo struct {
	DateOfMatch *int `xml:"DateOfMatch"json:"DateOfMatch,omitempty"`
	Sport struct {
		Id    *int    `xml:"id,attr"json:"id,omitempty"`
		Value *string `xml:",chardata"json:"value,omitempty"`
	} `xml:"Sport"json:"Sport,omitempty"`
	Category struct {
		Id    *int    `xml:"id,attr"json:"id,omitempty"`
		Value *string `xml:",chardata"json:"value,omitempty"`
	} `xml:"Category"json:"Category,omitempty"`
	Tournament struct {
		Id    *int    `xml:"id,attr"json:"id,omitempty"`
		Value *string `xml:",chardata"json:"value,omitempty"`
	} `xml:"Tournament"json:"Tournament,omitempty"`
	HomeTeam struct {
		Id       *int    `xml:"id,attr"json:"id,omitempty"`
		Value    *string `xml:",chardata"json:"value,omitempty"`
		Uniqueid *int    `xml:"uniqueid,attr"json:"uniqueid,omitempty"`
	} `xml:"HomeTeam"json:"HomeTeam,omitempty"`
	AwayTeam struct {
		Id       *int    `xml:"id,attr"json:"id,omitempty"`
		Value    *string `xml:",chardata"json:"value,omitempty"`
		Uniqueid *int    `xml:"uniqueid,attr"json:"uniqueid,omitempty"`
	} `xml:"AwayTeam"json:"AwayTeam,omitempty"`
	TvChannels []struct {
		TvChannel *string `xml:"TvChannel"json:"TvChannel,omitempty"`
	} `xml:"TvChannels"json:"TvChannels,omitempty"`
	Infos []struct {
		Type  string  `xml:"type,attr"json:"type,omitempty"`
		Value *string `xml:",chardata"json:"value,omitempty"`
	} `xml:"ExtraInfo>Info"json:"infos,omitempty"`
	//Streaming []struct {
	//	Channel struct {
	//		Id    *string`xml:"id,attr"json:"id,omitempty"`
	//		Value *string`xml:",chardata"json:"value,omitempty"`
	//	} `xml:"Channel"json:"Channel,omitempty"`
	//} `xml:"Streaming"json:"Streaming,omitempty"`
	//CoverpageInfo struct {
	//	Type struct {
	//		Id    *string`xml:"id,attr"json:"id,omitempty"`
	//		Value *string`xml:",chardata"json:"value,omitempty"`
	//	} `xml:"Type"json:"Type,omitempty"`
	//} `xml:"CoverpageInfo"json:"CoverpageInfo,omitempty"`
}

type Card struct {
	Canceled *string `xml:"canceled,attr"json:"canceled,omitempty"`
	Id       *int    `xml:"id,attr"json:"id,omitempty"`
	Player   *string `xml:"player,attr"json:"player,omitempty"`
	Team     *string `xml:"team,attr"json:"team,omitempty"`
	Time     *int    `xml:"time,attr"json:"time,omitempty"`
	Type     *string `xml:"type,attr"json:"type,omitempty"`
	Playerid *int    `xml:"playerid,attr"json:"playerid,omitempty"`
}
type Score struct {
	Away        *int    `xml:"away,attr"json:"away,omitempty"`
	Home        *int    `xml:"home,attr"json:"home,omitempty"`
	Id          *int    `xml:"id,attr"json:"id,omitempty"`
	Player      *string `xml:"player,attr"json:"player,omitempty"`
	Scoringteam *string `xml:"scoringteam,attr"json:"scoringteam,omitempty"`
	Time        *int    `xml:"time,attr"json:"time,omitempty"`
	Type        *string `xml:"type,attr"json:"type,omitempty"`
	Playerid    *int    `xml:"playerid,attr"json:"playerid,omitempty"`
}

type Odd struct {
	Id               *int        `xml:"id,attr"json:"id,omitempty"`
	Type             *string     `xml:"type,attr"json:"type,omitempty"`
	Subtype          *string     `xml:"subtype,attr"json:"subtype,omitempty"`
	Ftr              *string     `xml:"ftr,attr"json:"ftr,omitempty"`
	Specialoddsvalue *string     `xml:"specialoddsvalue,attr"json:"specialoddsvalue,omitempty"`
	Freetext         *string     `xml:"freetext,attr"json:"freetext,omitempty"`
	Active           *int        `xml:"active,attr"json:"active,omitempty"`
	Changed          *string     `xml:"changed,attr"json:"changed,omitempty"`
	Typeid           *int        `xml:"typeid,attr"json:"typeid,omitempty"`
	OddTypeId        *int        `xml:"oddtypeid,attr"json:"oddtypeid,omitempty"`
	Mostbalanced     *int        `xml:"mostbalanced,attr"json:"mostbalanced,omitempty"`
	OddsField        []OddsField `xml:"OddsField"json:"OddsField,omitempty"`
}

type OddsField struct {
	Type        *string  `xml:"type,attr"json:"type,omitempty"`
	Active      *int     `xml:"active,attr"json:"active,omitempty"`
	Typeid      *int     `xml:"typeid,attr"json:"typeid,omitempty"`
	OddsFieldId *int     `xml:"oddsfieldid,attr"json:"oddsfieldid,omitempty"`
	Outcome     *string  `xml:"outcome,attr"json:"outcome,omitempty"`
	Voidfactor  *float64 `xml:"voidfactor,attr"json:"voidfactor,omitempty"`
	Probability *string  `xml:"probability,attr"json:"probability,omitempty"`
	InnerValue  *string  `xml:",chardata"json:"value,omitempty"`
}
