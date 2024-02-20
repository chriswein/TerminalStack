package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Push(value string, db *sql.DB) {
	_, err := db.Exec(`INSERT INTO data
		(id, content)
		VALUES
		(NULL, ?)`, value)
	Check(err)
}

func Pop(db *sql.DB) string {
	rows, err := db.Query(`SELECT id,content FROM data ORDER BY id DESC LIMIT 1`)
	Check(err)
	var value string
	var id int
	rows.Next()
	err = rows.Scan(&id, &value)

	if err != nil {
		fmt.Println("")
		return ""
	}
	rows.Close()
	query := fmt.Sprintf(`DELETE FROM data WHERE id = %d`, id)
	result, err := db.Exec(query)
	if err != nil {
		fmt.Println(query, err.Error(), result)
	}
	return value
}
