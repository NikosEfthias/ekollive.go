package odd

import (
	"../../lib/db"
	"time"
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
	CreatedAt      time.Time `gorm:"column:createdAt"`
	UpdatedAt      time.Time `gorm:"column:updatedAt"`
}

func init() {
	Model = db.DB.Model(&Odd{})
	if !Model.HasTable(&Odd{}) {
		Model.CreateTable(&Odd{})
		Model.AddUniqueIndex("primary_key", "matchid", "oddid", "oddFieldTypeId", "oddTypeid")
	}
}
