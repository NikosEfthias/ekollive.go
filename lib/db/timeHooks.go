package db

import "time"

func (t *TimeFields) BeforeCreate() error {
	tm := time.Now()
	t.CreatedAt = &tm
	t.UpdatedAt = time.Now()
	return nil
}
func (t *TimeFields) BeforeUpdate() error {
	t.UpdatedAt = time.Now()
	return nil
}
