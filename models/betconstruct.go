package models

import (
	"encoding/json"
	"net"

	"github.com/k0kubun/pp"
	"github.com/mugsoft/tools"
)

type Fields struct {
	Id     int64  `json:"Id"`
	LangId string `json:"LangId"`
	Text   string `json:"Text"`
}

type LangFields struct {
	Id           int64               `json:"Id"`
	LangId       string              `json:"LangId"`
	Text         string              `json:"Text"`
	Name         string              `json:"Name"`
	Translations map[string][]Fields `json:"Translations"`
}

type Command struct {
	Command string                   `json:"Command"`
	Params  []map[string]interface{} `json:"Params,omitempty"`
	Timeout int                      `json:"__timeout,omitempty"`
	Objects []LangFields             `json:"Objects"`
}

func (l *Command) Send(sock net.Conn) error {
	d, err := json.Marshal(l)
	if nil != err {
		return err
	}
	pp.Println(string(d))
	ln := tools.Int2LE(uint32(len(d)))
	_, err = sock.Write(ln[:])
	if nil != err {
		return err
	}
	_, err = sock.Write(d)
	return err
}
