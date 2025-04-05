package restdb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Row struct {
	Id     int       `json:"id"`
	Date   time.Time `json:"date"`
	Close  float32   `json:"close"`
	Open   float32   `json:"open"`
	Volume int64     `json:"volume"`
	High   float32   `json:"high"`
	Low    float32   `json:"low"`
}

func Query(db *sql.DB, stmt string) []byte {
	// Select rows from existing table
	rows, err := db.Query(stmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Parse results of select query
	rowSlice := [](*Row){}
	for rows.Next() {
		row := Row{}
		switch err := rows.Scan(
			&row.Id, &row.Date, &row.Close,
			&row.Open, &row.Volume, &row.High,
			&row.Low,
		); err {
		case sql.ErrNoRows:
			fmt.Println("No rows returned")
		case nil:
			rowSlice = append(rowSlice, &row)
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
