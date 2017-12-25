package oddType

import "github.com/jinzhu/gorm"
import (
	"../../lib/db"
)

var Model *gorm.DB

type Oddtype struct {
	Oddtypeid    *int `gorm:"primary_key;not null"`
	Subtype      *string
	Type         *string
	Typeid       *int
	Oddtypevalue *string
	Status       *int     `gorm:"type:TINYINT;default:1"`
	MinStake     *float64 `gorm:"default:0;column:minStake"`
	MaxStake     *float64 `gorm:"default:0;column:maxStake"`
	MaxPay       *float64 `gorm:"default:0;column:maxPay"`
	ListOrder    *float64 `gorm:"default:0;column:listOrder"`
	db.TimeFields
}

func init() {
	Model = db.DB.Model(&Oddtype{})
	Model.AutoMigrate(&Oddtype{})
}
