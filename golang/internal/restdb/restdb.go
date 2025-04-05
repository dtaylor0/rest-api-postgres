package restdb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
)

type Row struct {
	Id     int    `json:"id"`
	Date   string `json:"date"`
	Close  string `json:"close"`
	Volume string `json:"volume"`
	High   string `json:"high"`
	Low    string `json:"low"`
}

func newRow(
	id int,
	date string,
	close string,
	volume string,
	high string,
	low string,
) *Row {
	r := Row{id, date, close, volume, high, low}
	return &r
}

func Query(db *sql.DB, stmt string) []byte {
	// Select rows from existing table
	rows, err := db.Query(stmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Parse results of select query
	var id int
	var date string
	var close string
	var volume string
	var high string
	var low string

	// result array
	rowSlice := [](*Row){}
	for rows.Next() {
		switch err := rows.Scan(&id, &date, &close, &volume, &high, &low); err {
		case sql.ErrNoRows:
			fmt.Println("No rows returned")
		case nil:
			r := newRow(id, date, close, volume, high, low)
			rowSlice = append(rowSlice, r)
		default:
			panic(err)
		}
	}

	res, err := json.Marshal(rowSlice)
	if err != nil {
		panic(err)
	}
	return res
}
