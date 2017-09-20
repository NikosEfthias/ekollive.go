package match

import "time"

func (od *Match) BeforeCreate() error {
	od.CreatedAt = time.Now()
	od.UpdatedAt = time.Now()
	return nil
}
func (od *Match) BeforeUpdate() error {
	od.UpdatedAt = time.Now()
	return nil
}
