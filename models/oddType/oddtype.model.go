package oddType

import "github.com/jinzhu/gorm"
import (
	"../../lib/db"
	"time"
)

var Model *gorm.DB

type Oddtype struct {
	Oddtypeid    *int `gorm:"primary_key;not null"`
	Subtype      *string
	Type         *string
	Typeid       *int
	Oddtypevalue *string
	UpdatedAt    time.Time    `gorm:"column:updatedAt"`
	CreatedAt    time.Time    `gorm:"column:createdAt"`
}

func init() {
	Model = db.DB.Model(&Oddtype{})
	Model.AutoMigrate(&Oddtype{})
}
