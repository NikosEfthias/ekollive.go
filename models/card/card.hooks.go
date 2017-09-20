package card

import "time"

func (od *Card) BeforeCreate() error {
	od.CreatedAt = time.Now()
	od.UpdatedAt = time.Now()
	return nil
}
func (od *Card) BeforeUpdate() error {
	od.UpdatedAt = time.Now()
	return nil
}
