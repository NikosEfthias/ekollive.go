package endpoints

type Success struct {
	Ok interface{} `json:"ok"`
}

type Error struct {
	Error interface{} `json:"error"`
}
