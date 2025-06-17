package models

import (
	"glasya-labolas/database"
)

type Entity struct {
	ID        int
	Nombre    string
	IP        string
	Hash      string
	CreatedAt string
}

func (e *Entity) Save() error {
	db := database.GetDB()
	_, err := db.Exec("INSERT INTO entities (nombre, ip, hash) VALUES (?, ?, ?)", e.Nombre, e.IP, e.Hash)
	return err
}

func ListEntities() ([]Entity, error) {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, nombre, ip, hash, created_at FROM entities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []Entity
	for rows.Next() {
		var e Entity
		rows.Scan(&e.ID, &e.Nombre, &e.IP, &e.Hash, &e.CreatedAt)
		entities = append(entities, e)
	}
	return entities, nil
}