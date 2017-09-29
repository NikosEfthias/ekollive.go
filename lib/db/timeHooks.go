package db

import "time"

func (t *TimeFields) BeforeCreate() error {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return nil
}
func (t *TimeFields) BeforeUpdate() error {
	t.UpdatedAt = time.Now()
	return nil
}
