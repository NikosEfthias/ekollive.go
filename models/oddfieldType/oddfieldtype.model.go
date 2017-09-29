package oddfieldType

import "github.com/jinzhu/gorm"
import (
	"../../lib/db"
)

type Oddfieldtype struct {
	Oddtypeid *int
	Typeid    *int
	Type      *string `gorm:"not null"`
	db.TimeFields
}

var Model *gorm.DB

func init() {
	Model = db.DB.Model(&Oddfieldtype{})
	if !Model.HasTable(&Oddfieldtype{}) {
		Model.CreateTable(&Oddfieldtype{})
		Model.AddUniqueIndex("primary_key", "oddtypeid", "typeid")
	}
}
