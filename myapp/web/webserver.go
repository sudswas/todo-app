package main

import (
	"encoding/json"
	"fmt"
	"log"
	db "myapp/db"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Record struct {
	Task  string `json:"task,omitempty"`
	Dueby string `json:"dueby,omitempty"`
}

func GetTodoAll(w http.ResponseWriter, r *http.Request) {
	db_handle := db.DBUtils{}
	recordList := db_handle.GetTodoAll("todo")
	json.NewEncoder(w).Encode(recordList)
}

func GetTodoByDate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db_handle := db.DBUtils{}
	dateSupplied := params["date"]
	if dateSupplied == "today" {
		dateToday := time.Now()
		dateSupplied = dateToday.Format("2006-01-02")
	}
	recordList := db_handle.GetTodoByDate("todo", dateSupplied)
	json.NewEncoder(w).Encode(recordList)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var record Record
	_ = json.NewDecoder(r.Body).Decode(&record)
	fmt.Println(record.Task)
	fmt.Println(record.Dueby)
	db_handle := db.DBUtils{}
	db_handle.Add("todo", record.Task, record.Dueby)
}

// our main function
func main() {
	// To be removed. We are sleeping here just to ensure that mariadb is pulled and up. We would ideally want to just use a liveness probe or something later
	time.Sleep(6 * time.Second)
	// Use caching of db_handle instead of initializing all the time.
	db_handle := db.DBUtils{}
	db_handle.Create("todo")
	router := mux.NewRouter()
	router.HandleFunc("/todo", GetTodoAll).Methods("GET")
	router.HandleFunc("/todo/{date}", GetTodoByDate).Methods("GET")
	router.HandleFunc("/todo", CreateTodo).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
