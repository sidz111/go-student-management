package repository

import (
	"database/sql"
	"strconv"

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

func (r *StudentRepository) GetallStudents() ([]model.Student, error) {
	query := "select * from student"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []model.Student
	for rows.Next() {
		var s model.Student
		err := rows.Scan(&s.Id, &s.Name, &s.Address)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}

func (r *StudentRepository) UpdateStudent(student model.Student) (string, error) {
	query := "update student set name = ?, address = ? where id = ?"
	_, err := r.DB.Exec(query, student.Name, student.Address, student.Id)
	if err != nil {
		return "Student Not Found with id " + strconv.Itoa(student.Id), err
	}
	return "Student Updated Successfull with id " + strconv.Itoa(student.Id), err
}
