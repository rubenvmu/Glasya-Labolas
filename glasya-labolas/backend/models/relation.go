package models

import (
	"glasya-labolas/database"
)

type Relation struct {
	ID        int
	SourceID  int
	TargetID  int
	Type      string
	CreatedAt string
}

func (r *Relation) Save() error {
	db := database.GetDB()
	_, err := db.Exec("INSERT INTO relations (source_id, target_id, type) VALUES (?, ?, ?)", r.SourceID, r.TargetID, r.Type)
	return err
}

func ListRelations() ([]Relation, error) {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, source_id, target_id, type, created_at FROM relations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relations []Relation
	for rows.Next() {
		var r Relation
		rows.Scan(&r.ID, &r.SourceID, &r.TargetID, &r.Type, &r.CreatedAt)
		relations = append(relations, r)
	}
	return relations, nil
}