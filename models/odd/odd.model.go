package odd

import (
	"../../lib/db"
	"github.com/jinzhu/gorm"
)

var Model *gorm.DB

type Odd struct {
	Oddid          *int
	Matchid        *int
	OddFieldTypeId *int `gorm:"column:oddFieldTypeId"`
	OddTypeId      *int `gorm:"column:oddTypeId"`
	Specialvalue   *string    `gorm:"column:specialvalue"`
	Mostbalanced   *int `gorm:"default:0"`
	Odd            *float64 `gorm:"default:1"`
	Active         *int `gorm:"not null;"`
	db.TimeFields
}

func init() {
	Model = db.DB.Model(&Odd{})
	if !Model.HasTable(&Odd{}) {
		Model.CreateTable(&Odd{})
		Model.AddUniqueIndex("primary_key", "matchid", "oddid", "oddFieldTypeId", "oddTypeid")
	}
}
