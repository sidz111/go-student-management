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
	}

	var student model.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid User", http.StatusBadRequest)
	}

	err = c.Repo.SaveStudent(student)

	if err != nil {
		http.Error(w, "Student Saved Failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Student Saved Done"))

}
