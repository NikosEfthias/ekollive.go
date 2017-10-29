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
	CreatedAt *time.Time `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}

var Model *gorm.DB

func init() {
	Model = db.DB.Model(&Oddfieldtype{})
	if !Model.HasTable(&Oddfieldtype{}) {
		Model.CreateTable(&Oddfieldtype{})
		Model.AddUniqueIndex("primary_key", "oddtypeid", "typeid")
	}
}
func (t *Oddfieldtype) BeforeCreate() error {
	tm := time.Now()
	t.CreatedAt = &tm
	t.UpdatedAt = time.Now()
	return nil
}
func (t *Oddfieldtype) BeforeUpdate() error {
	t.UpdatedAt = time.Now()
	return nil
}
