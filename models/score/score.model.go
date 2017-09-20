package score

import (
	"github.com/jinzhu/gorm"
	"time"
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
	CreatedAt   time.Time `gorm:"column:createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt"`
}

func init() {
	Model = db.DB.Model(&Score{})
	Model.AutoMigrate(&Score{})
}
