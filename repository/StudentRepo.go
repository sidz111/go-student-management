package repository

import (
	"database/sql"

	"github.com/sidz/student-management-go/model"
)

type StudentRepository struct {
	DB *sql.DB
}

func (r *StudentRepository) SaveStudent(s model.Student) error {
	query := "Insert into student(name, address) values(?,?)"
	_, err := r.DB.Exec(query, s.Name, s.Address)
	return err
}
