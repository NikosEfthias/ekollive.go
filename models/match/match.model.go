package match

import (
	"github.com/jinzhu/gorm"
	"../../lib/db"
	"time"
)

type Match struct {
	Matchid               *int `gorm:"primary_key"`
	Gamescore             *string `gorm:"column:gamescore"`
	Matchtime             *string `gorm:"column:matchtime"`
	Matchstatus           *string `gorm:"column:matchstatus"`
	Betstatus             *string `gorm:"column:betstatus"`
	Score                 *string `gorm:"column:score"`
	Service               *int `json:"service,omitempty"`
	Msgnumber             *int `gorm:"column:msgnumber"`
	Lastupdate            *string `gorm:"column:lastupdate"`
	Remainingtime         *string `json:"remainingtime,omitempty"`
	Suspendaway           *int `json:"suspendaway,omitempty"`
	Suspendhome           *int `json:"suspendhome,omitempty"`
	Tiebreak              *string `json:"tiebreak,omitempty"`
	Clockstop             *int `json:"clockstop,omitempty"`
	RemainingTimeinPeriod *string `json:"remainingtimeinperiod,omitempty"`
	Earlybetstatus        *string `gorm:"column:earlybetstatus"`
	Yrcardsaway           *int `gorm:"column:yellowredcardsaway"`
	Yrcardshome           *int `gorm:"column:yellowredcardshome"`
	Redcardsaway          *int `gorm:"column:redcardsaway"`
	Redcardshome          *int `gorm:"column:redcardshome"`
	Yellowcardsaway       *int `gorm:"column:yellowcardsaway"`
	Yellowcardshome       *int `gorm:"column:yellowcardshome"`
	Cornersaway           *int `gorm:"column:cornersaway"`
	Cornershome           *int `gorm:"column:cornershome"`
	Matchtimeextended     *string `gorm:"column:matchtimeextended"`
	Setscores             *string `gorm:"column:setscores"`
	Active                *int `gorm:"column:active;default:0"`
	CreatedAt             time.Time `gorm:"column:createdAt"`
	UpdatedAt             time.Time `gorm:"column:updatedAt"`
}

var Model *gorm.DB

func init() {
	Model = db.DB.Model(&Match{})
	Model.AutoMigrate(&Match{})
}
