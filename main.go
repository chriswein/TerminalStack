package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	pushorpop = "push" // Used for compilation of two distinct binarys
	file      = "terminalstack.sqlite"
)

func main() {

	db, err := sql.Open("sqlite3", file)

	var close func() = func() {
		err := db.Close()
		if err != nil {
			fmt.Printf("Could not close database: %s", err)
		}
	}

	defer close()
	Check(err)

	sqlStmt := "CREATE TABLE data (id INTEGER primary key, content TEXT);"
	_, createerr := db.Exec(sqlStmt)

	if createerr != nil {
		// The database was already created
	}

	if pushorpop != "pop" {
		stat, _ := os.Stdin.Stat()
		if stat.Mode()&os.ModeCharDevice == 0 {
			// Pipe input
			var stdin []byte
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				stdin = append(stdin, scanner.Bytes()...)
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
			Push(string(stdin), db)
		} else {
			// Use Program Arguments
			input := strings.Join(os.Args[1:], " ")
			Push(input, db)
		}
	} else {
		fmt.Println(Pop(db))
	}

}
