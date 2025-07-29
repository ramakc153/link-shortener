package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	connStr := "host=localhost port=5432 user=postgres password=admin dbname=link-shortener sslmode=disable"

	var err error

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	fmt.Println("DB connected")

}

func AddLink(key, long_url string) {
	short_url := fmt.Sprintf("http://localhost/%s", key)
	_, err := DB.Exec("INSERT INTO links (keys, long_url, short_url) VALUES ($1, $2, $3)", key, long_url, short_url)
	if err != nil {
		panic(err)
	}
}

func GetLink(key string) (*data, error) {
	var link_info data
	err := DB.QueryRow(
		"SELECT * FROM public.links WHERE keys=$1", key).Scan(
		&link_info.Key, &link_info.Long_url, &link_info.Short_url)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("data not found")
	}
	return &link_info, nil
}

func DeleteLink(key string) int64 {
	result, err := DB.Exec("DELETE FROM links WHERE keys=$1", key)

	if err != nil {
		panic(err)
	}
	rowaffected, _ := result.RowsAffected()
	return rowaffected
}
