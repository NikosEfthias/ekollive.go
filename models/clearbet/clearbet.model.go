package clearbet

import "time"
import (
	"../../lib/db"
	"github.com/jinzhu/gorm"
)

var Model *gorm.DB

type Clearbet struct {
	Matchid    *int `gorm:"column:matchId"`
	Oddid      *int `gorm:"column:oddId"`
	Outcome    *string `gorm:"column:outcome"`
	Type       *string `gorm:"column:type"`
	VoidFactor *float64 `gorm:"column:voidFactor"`
	Active     *int `gorm:"column:active"`
	CreatedAt  *time.Time `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"column:updatedAt"`
}

func init() {
	Model = db.DB.Model(&Clearbet{})
	if !Model.HasTable(&Clearbet{}) {
		Model.CreateTable(&Clearbet{})
		Model.AddUniqueIndex("primary_key", "matchId", "oddId", "type")
	}
}
func (t *Clearbet) BeforeCreate() error {
	tm := time.Now()
	t.CreatedAt = &tm
	t.UpdatedAt = time.Now()
	return nil
}
func (t *Clearbet) BeforeUpdate() error {
	t.UpdatedAt = time.Now()
	return nil
}
