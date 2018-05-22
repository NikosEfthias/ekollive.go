package oddType

import "github.com/jinzhu/gorm"
import (
	"../../lib/db"
)

var Model *gorm.DB

type Oddtype struct {
	Oddtypeid    *int     `gorm:"primary_key;not null" json:"oddtypeid"`
	Subtype      *string  `json:"subtype"`
	Type         *string  `json:"type"`
	Typeid       *int     `json:"typeid"`
	Sportid      *int     `json:"sportid"`
	Oddtypevalue *string  `json:"oddtypevalue"`
	HalfTime     *bool	  `gorm:"type:TINYINT;default:0;column:halftime"`
	Status       *int     `gorm:"type:TINYINT;default:1" json:"status"`
	MinStake     *float64 `gorm:"default:null;column:minStake" json:"min_stake"`
	BetLimit     *float64 `gorm:"default:null;column:betLimit" json:"bet_limit"`
	MaxStake     *float64 `gorm:"default:null;column:maxStake" json:"max_stake"`
	MaxPay       *float64 `gorm:"default:null;column:maxPay" json:"max_pay"`
	ListOrder    *float64 `gorm:"default:null;column:listOrder" json:"list_order"`
	db.TimeFields
}

func init() {
	Model = db.DB.Model(&Oddtype{})
	Model.AutoMigrate(&Oddtype{})
}
