package oddType

import "time"

func (od *Oddtype) BeforeCreate() error {
	od.CreatedAt = time.Now()
	od.UpdatedAt = time.Now()
	return nil
}
func (od *Oddtype) BeforeUpdate() error {
	od.UpdatedAt = time.Now()
	return nil
}
