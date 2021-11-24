package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
func ServeHome(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("<h1>Let's build awesome APIs 🔥🇮🇳</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type","application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses and find matching id and return response
	for _,course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
		return 
		}
	}
	json.NewEncoder(w).Encode("No course found with given id"+ params["id"])
	return 
}

func CreateOneCourse(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type","application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// if {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return 
	}

	// generate unique id 
	// convert id to string
	rand.Seed(time.Now().Unix())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}	


// dummy database
var courses [] Course

func main()  {
	fmt.Println("Getting into building API 🙏")


}