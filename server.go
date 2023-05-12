package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

type LessonData struct {
	TimeBegin time.Time `json:"TimeBegin"`
	TimeEnd   time.Time `json:"TimeEnd"`

	IsLecture bool `json:"IsLecture"`
	IsBreak   bool `json:"IsBreak"`
	IsNow     bool `json:"IsNow"`

	Name    string `json:"Name"`
	WhereIs int    `json:"WhereIs"`

	WhoTeaches       string `json:"WhoTeaches"`
	TeacherBio       string `json:"TeacherBio"`
	TeacherPhotoPath string `json:"TeacherPhotoPath"`
}

type DayData struct {
	Name   time.Weekday `json:"Name"`
	Number time.Time    `json:"Number"`

	IsMilitary bool `json:"IsMilitary"`
	IsHoliday  bool `json:"IsHoliday"`

	EvenWeekLessons    []LessonData `json:"EvenWeekLessons"`
	NotEvenWeekLessons []LessonData `json:"NotEvenWeekLessons"`
}

type WeekData struct {
	Number int  `json:"Number"`
	IsEven bool `json:"IsEven"`

	Week []DayData `json:"Week"`
}

type PageData struct {
	Title    string   `json:"Title"`
	Schedule WeekData `json:"Schedule"`
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

func lessonTimeFormatter(temp time.Time) string {
	return temp.Format("15:04")
}

func dayTimeFormatter(temp time.Time) string {
	return temp.Format("02.01.06")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	layoutFunctions := template.FuncMap{
		"lessonTimeFormatter": lessonTimeFormatter,
		"dayTimeFormatter":    dayTimeFormatter,
	}

	data := jsonParser()
	tmpl := template.Must(
		template.New("").Funcs(layoutFunctions).
			ParseFiles("templates/index_layout.html"))

	tmpl.ExecuteTemplate(w, "index_layout.html", data)
}

func main() {
	fileServer := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", homeHandler)

	fmt.Println("Listening port 1370...")
	if err := http.ListenAndServe(":1370", nil); err != nil {
		log.Fatal(err)
	}
}
