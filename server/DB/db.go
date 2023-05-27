package DB

import "database/sql"

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/chatbox_ai")
	if err != nil {
		panic(err)
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
