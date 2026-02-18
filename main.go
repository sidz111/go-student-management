package main

import (
	"fmt"
	"net/http"

	controller "github.com/sidz/student-management-go/Controller"
	"github.com/sidz/student-management-go/db"
	"github.com/sidz/student-management-go/repository"
)

func main() {
	fmt.Println("Student Management")
	db := db.ConnectDB()
	repo := &repository.StudentRepository{DB: db}
	controller := &controller.StudentController{Repo: repo}
	mux := http.NewServeMux()
	mux.HandleFunc("/student", controller.SaveStudent)
	mux.HandleFunc("/students", controller.GetAllStudents)
	mux.HandleFunc("/student/update", controller.UpdateStudent)
	mux.HandleFunc("/student/delete", controller.DeleteStudentById)
	mux.HandleFunc("/student/get", controller.GetStudentByID)
	fmt.Println("Server Stared at 8080")
	http.ListenAndServe(":8080", mux)
}
