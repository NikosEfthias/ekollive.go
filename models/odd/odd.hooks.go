package odd

import "time"

func (od *Odd) BeforeCreate() error {
	//if db.DB.NewRecord(od) {
	od.CreatedAt = time.Now()
	//}
	od.UpdatedAt = time.Now()
	return nil
}
func (od *Odd) BeforeUpdate() error {
	od.UpdatedAt = time.Now()
	return nil
}
