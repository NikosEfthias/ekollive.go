package odd

import (
	"../../lib/db"
	"github.com/jinzhu/gorm"
	"time"
)

var Model *gorm.DB

type Odd struct {
	Oddid          *int
	Matchid        *int
	OddFieldTypeId *int `gorm:"column:oddFieldTypeId"`
	OddTypeId      *int `gorm:"column:oddTypeId"`
	Specialvalue   *string    `gorm:"column:specialvalue"`
	Mostbalanced   *int `gorm:"default:0"`
	Odd            *float64 `gorm:"default:1"`
	Active         *int `gorm:"not null;"`
	CreatedAt      *time.Time `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt      time.Time `gorm:"column:updatedAt"`
}

func init() {
	Model = db.DB.Model(&Odd{}).Set("gorm:table_options", "ENGINE=MyISAM")
	if !Model.HasTable(&Odd{}) {
		Model.CreateTable(&Odd{})
		Model.AddUniqueIndex("primary_key", "matchid", "oddid", "oddFieldTypeId", "oddTypeid")
	}
}

func (t *Odd) BeforeCreate() error {
	tm := time.Now()
	t.CreatedAt = &tm
	t.UpdatedAt = time.Now()
	return nil
}
func (t *Odd) BeforeUpdate() error {
	t.UpdatedAt = time.Now()
	return nil
}
