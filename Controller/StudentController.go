package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func (c *StudentController) UpdateStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Invalid Method Type", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var student model.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid Student Data", http.StatusBadRequest)
		return
	}

	msg, err := c.Repo.UpdateStudent(student)
	if err != nil {
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": msg,
	})
}

func (c *StudentController) DeleteStudentById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid Method Request", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	msg, err := c.Repo.DeleteStudentById(id)

	if err != nil {
		http.Error(w, "Failed to Delete Student", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": msg,
	})
}
