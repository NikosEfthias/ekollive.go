package language

import (
	"../../lib/db"
	"github.com/jinzhu/gorm"
	"time"
)

var Model *gorm.DB

type Language struct {
	Id        *int
	DataType  *int `gorm:"column:dataType"`
	Lang      *string
	Key       *int
	Value     *string
	CreatedAt *time.Time `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt time.Time  `gorm:"column:updatedAt"`
}

func init() {
	Model = db.DB.Model(&Language{})
	if !Model.HasTable(&Language{}) {
		Model.CreateTable(&Language{})
		Model.AddUniqueIndex("primary_key", "matchid", "oddid", "oddFieldTypeId", "oddTypeid")
	}
}

func (t *Language) BeforeCreate() error {
	tm := time.Now()
	t.CreatedAt = &tm
	t.UpdatedAt = time.Now()
	return nil
}
func (t *Language) BeforeUpdate() error {
	t.UpdatedAt = time.Now()
	return nil
}
