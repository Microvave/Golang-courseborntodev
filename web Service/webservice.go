package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

func courseHandler(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/courses", courseHandler)
	http.ListenAndServe(":8080", nil)
}
