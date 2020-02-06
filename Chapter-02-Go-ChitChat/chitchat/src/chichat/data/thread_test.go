package data

//
import (
	"database/sql"
	"log"
	"testing"
)

func db() (DbTmp *sql.DB)  {
	var err error
	DbTmp, err = sql.Open("postgres", "dbname=chitchat user=postgres password=123456 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return DbTmp
}

// Delete all threads from database
func ThreadDeleteAll() (err error) {
	db := db()
	defer db.Close()
	statement := "delete from threads"
	_, err = db.Exec(statement)
	if err != nil {
		return
	}
	return
}

func Test_CreateThread(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	conv, err := users[0].CreateThread("My first thread")
	if err != nil {
		t.Error(err, "Cannot create thread")
	}
	if conv.UserId != users[0].Id {
		t.Error("User not linked with thread")
	}
}
