package main

import "database/sql"
import _ "github.com/Go-SQL-Driver/MySQL"

func RecordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()
/*
	if _, err = tx.Exec("CREATE  table me"); err != nil {
		return
	}*/
	if _, err = tx.Exec("INSERT INTO me (id, name) VALUES (?, ?)", userID, productID); err != nil {
		return
	}
	return
}

func main() {
	// @NOTE: the real connection is not required for tests
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mysql_test?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = RecordStats(db, 1 /*some user id*/, 5 /*some product id*/); err != nil {
		panic(err)
	}
}
