package oddfieldType

import "github.com/jinzhu/gorm"
import (
	"../../lib/db"
	"time"
)

type Oddfieldtype struct {
	Oddtypeid *int
	Typeid    *int
	Type      *string `gorm:"not null"`
	UpdatedAt time.Time    `gorm:"column:updatedAt"`
	CreatedAt time.Time    `gorm:"column:createdAt"`
}

var Model *gorm.DB

func init() {
	Model = db.DB.Model(&Oddfieldtype{})
	if !Model.HasTable(&Oddfieldtype{}) {
		Model.CreateTable(&Oddfieldtype{})
		Model.AddUniqueIndex("primary_key", "oddtypeid", "typeid")
	}
}
