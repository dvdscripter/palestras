package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const schema = `
create table users(
	name TEXT
);
insert into users values("David");
insert into users values("Lilian");
`

type DB struct {
	*sql.DB
}

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	myDB := &DB{db}

	if _, err := db.Exec(schema); err != nil {
		log.Fatal(err)
	}

	id, err := myDB.GetUserByName("Lilian")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
}

func (db *DB) GetUserByName(name string) (int, error) {

	var id int

	query := "select oid from users where name = '" + name + "'"

	err := db.QueryRow(query).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil

}
