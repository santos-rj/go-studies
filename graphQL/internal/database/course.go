package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name string, description string, categoryId string) (*Course, error) {
	query := `INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4) RETURNING id`
	err := c.db.QueryRow(query, uuid.New().String(), name, description, categoryId).Scan(&c.ID)
	if err != nil {
		return nil, err
	}
	return c, nil
}
