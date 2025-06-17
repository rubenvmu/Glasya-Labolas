package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./glasya.db")
	if err != nil {
		log.Fatal("❌ Error al abrir la base de datos:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("❌ Error al conectar con la base de datos:", err)
	}

	createTables()
}

func GetDB() *sql.DB {
	return db
}

func createTables() {
	entityTable := `
	CREATE TABLE IF NOT EXISTS entities (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		ip TEXT,
		hash TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	relationTable := `
	CREATE TABLE IF NOT EXISTS relations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		source_id INTEGER,
		target_id INTEGER,
		relation TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(entityTable); err != nil {
		log.Fatal("❌ Error creando tabla entities:", err)
	}

	if _, err := db.Exec(relationTable); err != nil {
		log.Fatal("❌ Error creando tabla relations:", err)
	}
}