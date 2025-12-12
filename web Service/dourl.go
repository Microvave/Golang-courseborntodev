package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Course struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tel  string `json:"tel"`
}

var CourseList []Course

func init() {
	CourseJson := `[
  {
    "ID": 1,
    "Name": "Nutasit",
    "Tel": "0812345678"
  },
  {
    "ID": 2,
    "Name": "Wave",
    "Tel": "0899998888"
  },
  {
    "ID": 3,
    "Name": "Microwave",
    "Tel": "0905554444"
  }
]`
	err := json.Unmarshal([]byte(CourseJson), &CourseList)
	if err != nil {
		log.Fatal(err)
	}
}

func getnextID() int {
	highestID := -1
	for _, course := range CourseList {
		if course.ID > highestID {
			highestID = course.ID
		}
	}
	return highestID + 1
}

func findID(ID int) (*Course, int) {
	for i, course := range CourseList {
		if course.ID == ID {
			return &course, i
		}
	}
	return nil, 0
}
func courseHandler(w http.ResponseWriter, r *http.Request) {
	// URL: /course/{id}
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	// parts = ["course", "4"]

	if len(parts) != 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ID, err := strconv.Atoi(parts[1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	course, index := findID(ID)
	if course == nil {
		http.Error(w, fmt.Sprintf("no course %d", ID), http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(course)

	case http.MethodPut:
		var updateCourse Course
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := json.Unmarshal(body, &updateCourse); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if updateCourse.ID != ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		CourseList[index] = updateCourse
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	CourseJson, err := json.Marshal(CourseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(CourseJson)
	case http.MethodPost:
		var newCourse Course
		Bodybyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(Bodybyte, &newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newCourse.ID != 0 {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		newCourse.ID = getnextID()
		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}

}
func main() {
	http.HandleFunc("/course/", courseHandler) // <--- สำคัญ!!
	http.HandleFunc("/courses", coursesHandler)
	http.ListenAndServe(":8080", nil)
}
