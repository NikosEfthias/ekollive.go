package score

import (
	"github.com/jinzhu/gorm"
	"../../lib/db"
)

var Model *gorm.DB

type Score struct {
	Scoreid     *int `gorm:"column:scoreid;primary_key"`
	Home        *int `gorm:"column:home"`
	Away        *int `gorm:"column:away"`
	Player      *string `gorm:"column:player"`
	Scoringteam *string `gorm:"column:scoringteam"`
	Scoretime   *int `gorm:"column:scoretime"`
	Scoretype   *string `gorm:"column:scoretype"`
	Playerid    *int `gorm:"column:playerid"`
	Matchid     *int `gorm:"column:matchid"`
	Matchtime   *string `gorm:"column:matchtime"`
	Matchscore  *string `gorm:"column:matchscore"`
	Matchstatus *string `gorm:"column:matchstatus"`
	db.TimeFields
}

func (scr *Score) TableName() string {
	return "Scores"
}
func init() {
	Model = db.DB.Model(&Score{})
	Model.AutoMigrate(&Score{})
}
