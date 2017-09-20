package origin

import "github.com/jinzhu/gorm"
import (
	"../../../lib/db"
)

var Model *gorm.DB

type Allow struct {
	ID       string `gorm:"column:id;primary_key"`
	Password string `gorm:"default:null"json:"-"`
}

func (a *Allow) TableName() string {
	return "Origins"
}

func init() {
	Model = db.DB.Model(&Allow{})
	Model.AutoMigrate(&Allow{})
}
func CheckOk(org string, passwd ...string) bool {
	_ = passwd //not implemented yet
	out := new(Allow)
	Model.Where(&Allow{ID: org}).First(out)
	if out.ID != org {
		return false
	}
	return true
}
