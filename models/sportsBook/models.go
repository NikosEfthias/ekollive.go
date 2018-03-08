package sportsBook

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Sport struct {
	SportId   *int    `gorm:"column:sportId"`
	SportName *string `gorm:"column:sportName"`
	Lang      string  `gorm:"lang;default:en"`
}

func (spr Sport) Tablename() string {
	return "sport"
}

type Category struct {
	SportId      *int    `gorm:"column:sportId"`
	Categoryid   *int    `gorm:"column:categoryId"`
	CategoryName *string `gorm:"column:categoryName"`
	Lang         string  `gorm:"column:lang;default:en"`
}

func (c Category) Tablename() string {
	return "category"
}

type Tournament struct {
	SportId        *int    `gorm:"column:sportId"`
	Categoryid     *int    `gorm:"column:categoryId"`
	TournamentId   *int    `gorm:"column:tournamentId"`
	TournamentName *string `gorm:"column:tournamentName"`
	Lang           string  `gorm:"column:lang;default:en"`
}

func (c Tournament) Tablename() string {
	return "tournament"
}

type Competitor struct {
	CompId       *int    `gorm:"column:compId"`
	Comp2Id      *int    `gorm:"column:CompId2"`
	SportId      *int    `gorm:"column:sportId"`
	Categoryid   *int    `gorm:"column:categoryId"`
	TournamentId *int    `gorm:"column:tournamentId"`
	CompName     *string `gorm:"column:compName"`
	Lang         string  `gorm:"column:lang;default:en"`
}

func (c Competitor) Tablename() string {
	return "competitor"
}

type Match struct {
	SportId      *int   `gorm:"column:sportId"`
	Categoryid   *int   `gorm:"column:categoryId"`
	TournamentId *int   `gorm:"column:tournamentId"`
	Matchid      *int   `gorm:"column:matchId"`
	Comp1        *int   `gorm:"column:comp1"`
	Comp2        *int   `gorm:"column:comp2"`
	Matchdate    string `gorm:"column:matchDate"`
	PeriodLength *int   `gorm:"column:periodLength"`
	LiveActive   string `gorm:"column:liveActive"`
}

func (m Match) Tablename() string {
	return "match"
}
func (m Match) Update(db *sql.DB) {
	upsert(db, m.Tablename(), &m, &Match{LiveActive: "1", Matchdate: m.Matchdate, PeriodLength: m.PeriodLength})
}
func upsert(db *sql.DB, tableName string, doc *Match, update *Match) {
	var updateStr = "insert into `" + tableName + "` (?) Values (?) on duplicate key update liveActive='" +
		update.LiveActive + "' ,matchDate='" + update.Matchdate + "'"
	if update.PeriodLength != nil {
		updateStr += ",periodLength=" + strconv.Itoa(*update.PeriodLength)
	}
	var fields = ""
	var values = ""
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
		switch field.Interface().(type) {
		case time.Time:
			val := field.Interface().(time.Time)
			if val == (time.Time{}) || name == "updatedAt" {
				val = time.Now().UTC()
			}
			values += "'" + val.Format("2006-01-02 15:04:05") + "',"

		case string:
			values += "'" + field.Interface().(string) + "',"

		default:
			values += fmt.Sprintf("%v,", field.Interface())
		}
	}
	updateStr = strings.Replace(updateStr, "?", fields[:len(fields)-1], 1)
	updateStr = strings.Replace(updateStr, "?", values[:len(values)-1], 1)
	//fmt.Println("\x1B[32m", updateStr, "\x1B[0m")
	db.Exec(updateStr)

}
