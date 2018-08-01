package models

type BetconstructData struct {
	Command *string
	Objects []Object
	Error   *map[string]interface{}
	Type    *string
}
type Object map[string]interface{}
