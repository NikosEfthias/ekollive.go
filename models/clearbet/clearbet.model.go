package clearbet

import "time"
import (
	"../../lib/db"
	"github.com/jinzhu/gorm"
)

var Model *gorm.DB

type Clearbet struct {
	Id        int `gorm:"column:id;primary_key;"`
	Matchid   *int `gorm:"column:matchId"`
	Oddid     *int `gorm:"column:oddId"`
	Starttime *time.Time `gorm:"column:startTime"`
	Endtime   *time.Time `gorm:"column:endTime"`
	Xmltime   *time.Time `gorm:"column:xmlTime"`
	db.TimeFields
}

func init() {
	Model = db.DB.Model(&Clearbet{})
	Model.AutoMigrate(&Clearbet{})
}
