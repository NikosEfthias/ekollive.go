package oddfieldType

import "github.com/jinzhu/gorm"
import (
	"../../lib/db"
	"time"
)

type Oddfieldtype struct {
	Oddtypeid *int
	Typeid    *int
	Type      *string    `gorm:"not null"`
	MinStake  *float64   `gorm:"default:0;column:minStake"`
	MaxStake  *float64   `gorm:"default:0;column:maxStake"`
	MaxPay    *float64   `gorm:"default:0;column:maxPay"`
	ListOrder *float64   `gorm:"default:0;column:listOrder"`
	CreatedAt *time.Time `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt time.Time  `gorm:"column:updatedAt"`
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
