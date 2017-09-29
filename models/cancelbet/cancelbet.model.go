package cancelbet

import (
	"github.com/jinzhu/gorm"
	"time"
	"../../lib/db"
)

var Model *gorm.DB

type Cancelbet struct {
	Id        int `gorm:"column:id;primary_key;"`
	Matchid   *int `gorm:"column:matchId"`
	Oddid     *int `gorm:"column:oddId"`
	Starttime *time.Time `gorm:"column:startTime"`
	Endtime   *time.Time `gorm:"column:endTime"`
	Xmltime   *time.Time `gorm:"column:xmlTime"`
	db.TimeFields
}

func init() {
	Model = db.DB.Model(&Cancelbet{})
	Model.AutoMigrate(&Cancelbet{})
}
