package controllers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/pp"
	"github.com/mugsoft/ekollive.go/lib"
	"github.com/mugsoft/ekollive.go/models"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", *lib.DB)
err:
	if nil != err {
		panic(err)
	}
	err = db.Ping()
	if nil != err {
		goto err
	}
}

func Add_me_if_you_can(c *models.Command) {
	switch c.Command {
	case "GetTranslations":
		pp.Println("trans")
		go handle__translations(c)
	case "GetLanguages":
		pp.Println("lang")
		go handle__languages(c)
	case "HeartBeat":
		pp.Println(c.Command)
	}
}
func handle__languages(c *models.Command) {
	var q = "insert ignore into languages (langId,id,text) Values (?,?,?)"
	s, err := db.Prepare(q)
	if nil != err {
		fmt.Fprintln(os.Stderr, "handle__translations:", err.Error())
		return
	}
	defer s.Close()
	for _, o := range c.Objects {
		_, err = s.Exec(o.LangId, o.Id, o.Name)
		if nil != err {
			fmt.Fprintln(os.Stderr, "handle__languages:", err.Error())
			continue
		}
	}
}
func handle__translations(c *models.Command) {
	var q = "INSERT IGNORE INTO `translation` (`langId`, `name`, `nameId`) VALUES (?,?,?)"
	s, err := db.Prepare(q)
	if nil != err {
		fmt.Fprintln(os.Stderr, "handle__translations:", err.Error())
		return
	}
	defer s.Close()
	for _, obj := range c.Objects {
		for _, v := range obj.Translations {
			for _, translation := range v {
				_, err = s.Exec(translation.LangId, translation.Text, translation.Id)
				if nil != err {
					fmt.Fprintln(os.Stderr, "handle__translations::", err.Error())
					continue
				}
			}
		}
	}
}
