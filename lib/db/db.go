package db

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"../../lib"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var DB2 *gorm.DB

type TimeFields struct {
	CreatedAt *time.Time `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt time.Time  `gorm:"column:updatedAt"`
}

func checkPingDBS(db1, db2 *sql.DB) {
	var c1, c2 = make(chan bool), make(chan bool)
	go func(c1, c2 chan bool) {
		go func() {
			for {
				<-time.After(time.Second)
				err := db1.Ping()
				if nil != err {
					c1 <- true
					continue
				}
				c1 <- false
			}
		}()
		go func() {
			for {
				<-time.After(time.Second)
				err := db2.Ping()
				if nil != err {
					c2 <- true
					continue
				}
				c2 <- false
			}
		}()
	}(c1, c2)
	for {
		select {
		case d := <-c1:
			if d {
				fmt.Println("the first db doesnt respond")
				os.Exit(1)
			}
		case <-time.After(time.Second * 3):
			fmt.Println("the first db doesnt respond")
			os.Exit(1)
		}
		select {
		case d := <-c2:
			if d {
				fmt.Println("the second db doesnt respond")
				os.Exit(1)
			}
		case <-time.After(time.Second * 3):
			fmt.Println("the second db doesnt respond")
			os.Exit(1)
		}
	}
}

var queryWriter io.Writer

func init() {
	var err error
	DB, err = gorm.Open("mysql", (*lib.DB)+"?parseTime=true")
	if nil != err {
		log.Fatalln(err)
	}
	err = DB.DB().Ping()
	if nil != err {
		log.Fatalln(err)
	}
	DB.DB().SetMaxIdleConns(4)
	DB.DB().SetMaxOpenConns(100)
	DB2, err = gorm.Open("mysql", (*lib.DB2)+"?parseTime=true")
	if nil != err {
		log.Fatalln(err)
	}
	err = DB2.DB().Ping()
	if nil != err {
		log.Fatalln(err)
	}
	if os.Getenv("PrintUpserts") == "print" {
		queryWriter, err = os.OpenFile("queries.dump", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
		if nil != err {
			log.Fatalln(err)
		}
	}
	DB2.DB().SetMaxIdleConns(4)
	DB.DB().SetMaxOpenConns(100)
	//DB.LogMode(false)
}
func Upsert(db *sql.DB, tableName string, doc interface{}) {
	var updateStr = "insert into `" + tableName + "` (?) Values (?) on duplicate key update ?"
	var fields = ""
	var values = ""
	var duplicate = ""
	var q = reflect.ValueOf(doc)
	if q.Kind() == reflect.Ptr {
		q = q.Elem()
	}

	for i := 0; i < q.NumField(); i++ {
		if reflect.Ptr == q.Field(i).Kind() && q.Field(i).IsNil() {
			continue
		}
		name := strings.ToLower(q.Type().Field(i).Name)
		if tag := q.Type().Field(i).Tag.Get("gorm"); tag != "" && strings.Contains(tag, "column") {
			spl := strings.Split(tag, ";")
			for _, part := range spl {
				if !strings.Contains(part, "column") {
					continue
				}
				name = strings.Split(part, ":")[1]
			}
		}
		if q.Field(i).Type() == reflect.TypeOf(time.Time{}) && name != "updatedAt" {
			continue
		}
		fields += "`" + name + "`,"
		var field reflect.Value
		if q.Field(i).Kind() == reflect.Ptr {
			field = q.Field(i).Elem()
		} else {
			field = q.Field(i)
		}
		duplicate += name + "="
		switch field.Interface().(type) {
		case time.Time:
			val := field.Interface().(time.Time)
			if val == (time.Time{}) || name == "updatedAt" {
				val = time.Now().UTC()
			}
			values += "'" + val.Format("2006-01-02 15:04:05") + "',"
			duplicate += "'" + val.Format("2006-01-02 15:04:05") + "',"

		case string:
			values += "'" + field.Interface().(string) + "',"
			duplicate += "'" + field.Interface().(string) + "',"

		default:
			values += fmt.Sprintf("%v,", field.Interface())
			duplicate += fmt.Sprintf("%v,", field.Interface())
		}
	}
	updateStr = strings.Replace(updateStr, "?", fields[:len(fields)-1], 1)
	updateStr = strings.Replace(updateStr, "?", values[:len(values)-1], 1)
	updateStr = strings.Replace(updateStr, "?", duplicate[:len(duplicate)-1], 1)
	if os.Getenv("PrintUpserts") == "print" {
		fmt.Fprintln(queryWriter, updateStr)
	}
	//fmt.Println("\x1B[31m", updateStr, "x1B[0m")
	db.Exec(updateStr)

}
