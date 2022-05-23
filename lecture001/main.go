package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

type Student struct {
	Name     string
	Score    float64
	IsPassed bool
	Ranking  string
}

func NewStudent(name string, score float64) Student {
	var std = Student{
		Name:  name,
		Score: score,
	}
	std = GetStudentRanking(std)
	return std
}

func GetStudentRanking(std Student) Student {
	if std.Score >= 90 && std.Score <= 100 {
		std.Ranking = "A"
		std.IsPassed = true
	} else if std.Score >= 80 && std.Score < 90 {
		std.Ranking = "B"
		std.IsPassed = true
	} else if std.Score >= 70 && std.Score < 80 {
		std.Ranking = "C"
		std.IsPassed = true
	} else {
		std.Ranking = "F"
		std.IsPassed = false
	}
	return std
}

var database []Student = []Student{
	NewStudent("Ahmad", 84.5),
	NewStudent("Karim", 99.9),
}

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		temp, err := template.ParseFiles("./index.html")
		if request.Method == "POST" {
			request.ParseForm()

			var studentName string = request.PostForm.Get("name")
			studentScore, err := strconv.ParseFloat(request.PostForm.Get("score"), 64)
			if err != nil {
				http.Redirect(response, request, "/", 302)
				return
			}

			var newStudent = NewStudent(studentName, studentScore)
			database = append(database, newStudent)
		}
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			return
		}
		temp.Execute(response, database)

	})

	http.HandleFunc("/rest", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		resp, err := json.Marshal(database)
		if err != nil {
			return
		}
		response.Write(resp)
	})

	http.ListenAndServe(":8080", nil)

}
