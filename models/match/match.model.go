package match

import (
	"time"

	"../../lib/db"
	"github.com/jinzhu/gorm"
)

type Match struct {
	Matchid               *int       `gorm:"primary_key" json:"matchid"`
	SportId               *int       `gorm:"column:sportid" json:"sport_id"`
	CategoryId            *int       `gorm:"column:categoryid" json:"category_id"`
	TournamentId          *int       `gorm:"column:tournamentid" json:"tournament_id"`
	Gamescore             *string    `gorm:"column:gamescore" json:"gamescore"`
	Matchtime             *string    `gorm:"column:matchtime" json:"matchtime"`
	Matchstatus           *string    `gorm:"column:matchstatus" json:"matchstatus"`
	Betstatus             *string    `gorm:"column:betstatus" json:"betstatus"`
	Score                 *string    `gorm:"column:score" json:"score"`
	Service               *int       `json:"service,omitempty"`
	Msgnumber             *int       `gorm:"column:msgnumber" json:"msgnumber"`
	Lastupdate            *string    `gorm:"column:lastupdate" json:"lastupdate"`
	Remainingtime         *string    `json:"remainingtime,omitempty"`
	Suspendaway           *int       `json:"suspendaway,omitempty"`
	Suspendhome           *int       `json:"suspendhome,omitempty"`
	Tiebreak              *string    `json:"tiebreak,omitempty"`
	Clockstop             *int       `json:"clockstop,omitempty"`
	RemainingTimeinPeriod *string    `json:"remainingtimeinperiod,omitempty"`
	Earlybetstatus        *string    `gorm:"column:earlybetstatus" json:"earlybetstatus"`
	Yrcardsaway           *int       `gorm:"column:yellowredcardsaway" json:"yrcardsaway"`
	Yrcardshome           *int       `gorm:"column:yellowredcardshome" json:"yrcardshome"`
	Redcardsaway          *int       `gorm:"column:redcardsaway" json:"redcardsaway"`
	Redcardshome          *int       `gorm:"column:redcardshome" json:"redcardshome"`
	Yellowcardsaway       *int       `gorm:"column:yellowcardsaway" json:"yellowcardsaway"`
	Yellowcardshome       *int       `gorm:"column:yellowcardshome" json:"yellowcardshome"`
	Cornersaway           *int       `gorm:"column:cornersaway" json:"cornersaway"`
	Cornershome           *int       `gorm:"column:cornershome" json:"cornershome"`
	Matchtimeextended     *string    `gorm:"column:matchtimeextended" json:"matchtimeextended"`
	Setscores             *string    `gorm:"column:setscores" json:"setscores"`
	Active                *int       `gorm:"column:active;default:0" json:"active"`
	CreatedAt             *time.Time `gorm:"column:createdAt;default:current_timestamp" json:"created_at"`
	UpdatedAt             time.Time  `gorm:"column:updatedAt" json:"updated_at"`
}

var Model *gorm.DB

func init() {
	Model = db.DB.Model(&Match{})
	Model.AutoMigrate(&Match{})
}

func (t *Match) BeforeCreate() error {
	tm := time.Now()
	t.CreatedAt = &tm
	t.UpdatedAt = time.Now()
	return nil
}
func (t *Match) BeforeUpdate() error {
	t.UpdatedAt = time.Now()
	return nil
}
