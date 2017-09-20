package card

import "github.com/jinzhu/gorm"
import (
	"../../lib/db"
	"time"
)

var Model *gorm.DB

type Card struct {
	Cardid      *int `gorm:"column:cardid;primary_key"`
	Canceled    *string `gorm:"column:canceled"`
	Player      *string `gorm:"column:player"`
	Cardteam    *string `gorm:"column:cardteam"`
	Cardtime    *int `gorm:"column:cardtime"`
	Cardtype    *string `gorm:"column:cardtype"`
	Playerid    *int `gorm:"column:playerid"`
	Matchid     *int `gorm:"column:matchid"`
	Matchtime   *string `gorm:"column:matchtime"`
	Matchscore  *string `gorm:"column:matchscore"`
	Matchstatus *string `gorm:"column:matchstatus"`
	CreatedAt   time.Time `gorm:"column:createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt"`
}

func init() {
	Model = db.DB.Model(&Card{})
	Model.AutoMigrate(&Card{})
}
