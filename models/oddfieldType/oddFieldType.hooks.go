package oddfieldType

import "time"

func (od *Oddfieldtype) BeforeCreate() error {
	//if db.DB.NewRecord(od) {
	od.CreatedAt = time.Now()
	//}
	od.UpdatedAt = time.Now()
	return nil
}
func (od *Oddfieldtype) BeforeUpdate() error {
	od.UpdatedAt = time.Now()
	return nil
}
