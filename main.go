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
	http.HandleFunc("/student", controller.SaveStudent)
	http.HandleFunc("/students", controller.GetAllStudents)
	fmt.Println("Server Stared at 8080")
	http.ListenAndServe(":8080", nil)
}
