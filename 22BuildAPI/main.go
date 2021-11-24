package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// model for course
type Course struct {
	CourseId  string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// middlewares
func (c *Course ) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// controllers
func serveHome(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("<h1>Let's build awesome APIs 🔥🇮🇳</h1>"))
}

func getAllCourses(w http.ResponseWriter, r http.Request)  {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

// dummy database
var courses [] Course

func main()  {
	fmt.Println("Getting into building API 🙏")


}