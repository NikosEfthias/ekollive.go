package score

import "time"

func (scr *Score) TableName() string {
	return "Scores"
}
func (od *Score) BeforeCreate() error {
	od.CreatedAt = time.Now()
	od.UpdatedAt = time.Now()
	return nil
}
func (od *Score) BeforeUpdate() error {
	od.UpdatedAt = time.Now()
	return nil
}
