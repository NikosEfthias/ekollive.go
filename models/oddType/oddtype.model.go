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
	db.TimeFields
}

func init() {
	Model = db.DB.Model(&Oddtype{})
	Model.AutoMigrate(&Oddtype{})
}
