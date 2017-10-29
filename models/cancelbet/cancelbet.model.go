package cancelbet

import (
	"github.com/jinzhu/gorm"
	"../../lib/db"
	"time"
)

var Model *gorm.DB

type Cancelbet struct {
	Matchid   *int `gorm:"column:matchId"`
	Oddid     *int `gorm:"column:oddId"`
	Starttime *int `gorm:"column:startTime;type:BIGINT"`
	Endtime   *int `gorm:"column:endTime;type:BIGINT"`
	CreatedAt *time.Time `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}

func init() {
	Model = db.DB.Model(&Cancelbet{})
	if !Model.HasTable(&Cancelbet{}) {
		Model.CreateTable(&Cancelbet{})
		Model.AddUniqueIndex("primary_key", "matchId", "oddId")
	}
}

func (t *Cancelbet) BeforeCreate() error {
	tm := time.Now()
	t.CreatedAt = &tm
	t.UpdatedAt = time.Now()
	return nil
}
func (t *Cancelbet) BeforeUpdate() error {
	t.UpdatedAt = time.Now()
	return nil
}
