package controller

import (
	"encoding/json"
	"net/http"

	"github.com/sidz/student-management-go/model"
	"github.com/sidz/student-management-go/repository"
)

type StudentController struct {
	Repo *repository.StudentRepository
}

func (c *StudentController) SaveStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var student model.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid User", http.StatusBadRequest)
		return
	}

	err = c.Repo.SaveStudent(student)

	if err != nil {
		http.Error(w, "Student Saved Failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Student Saved Done"))

}

func (c *StudentController) GetAllStudents(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Method type", http.StatusMethodNotAllowed)
		return
	}
	students, err := c.Repo.GetallStudents()
	if err != nil {
		http.Error(w, "Failsed to fetch Students", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
