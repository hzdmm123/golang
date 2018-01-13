package main

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestRecordStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("mock error: '%s' ", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	//mock.ExpectExec("UPDATE me").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO me values (id,name) (?,?)").WithArgs(3,4)
	mock.ExpectCommit()

	if err = RecordStats(db, 4, 3); err != nil {
		t.Errorf("exe error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("not implements: %s", err)
	}
}
