package filters

import "github.com/jinzhu/gorm"
import "../../lib/db"

type Filter struct {
	Matchid *string `gorm:"column:matchid;not null"`
	Key     *string `gorm:"column:key;not null"`
	Filter  *string `gorm:"column:filter;not null"`
}

var Model *gorm.DB

func init() {
	Model = db.DB.Model(&Filter{})
	if !Model.HasTable(&Filter{}) {
		Model.CreateTable(&Filter{})
		Model.AddUniqueIndex("primary_key", "matchid", "key")
	}
}
