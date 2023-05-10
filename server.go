package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type LessonData struct {
	TimeBegin time.Time `json:"TimeBegin"`
	TimeEnd   time.Time `json:"TimeEnd"`

	IsBreak bool `json:"IsBreak"`
	IsNow   bool `json:"IsNow"`

	Name    string `json:"Name"`
	WhereIs int    `json:"WhereIs"`

	WhoTeaches string `json:"WhoTeaches"`
	TeacherBio string `json:"TeacherBio"`
}

type DayData struct {
	Name   time.Weekday `json:"Name"`
	Number time.Time    `json:"Number"`

	IsMilitary bool `json:"IsMilitary"`

	Lessons []LessonData `json:"Lessons"`
}

type ScheduleData struct {
	IsEvenWeek bool      `json:"IsEvenWeek"`
	Week       []DayData `json:"Week"`
}

type PageData struct {
	Title    string       `json:"Title"`
	Schedule ScheduleData `json:"Schedule"`
}

func jsonParser() (output PageData) {
	input, err := ioutil.ReadFile("database/schedule.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(input, &output); err != nil {
		log.Fatal(output)
	}

	return
}

func tableMaker(db PageData) (table string) {

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Listening port 1370...")
	if err := http.ListenAndServe(":1370", nil); err != nil {
		log.Fatal(err)
	}
}
