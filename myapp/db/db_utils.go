package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type record struct {
	Task   string
	Dueby  string
	Status string
}

type DBUtils struct{}

func (d DBUtils) Create(name string) {
	fmt.Println("Going to create the Database")
	// read the password from the env variable
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
		fmt.Println("There was a problem connecting to the db", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todo ( id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, task varchar(32), dueby DATE, status boolean not null default 0)")
	if err != nil {
		panic(err)
	}
}

func (d DBUtils) Add(db_name string, task string, dueby string) bool {
	// use connection pools and read the password from env variable.
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
		fmt.Println("There was a problem connecting to the db", err)
	}

	defer db.Close()

	_, err = db.Exec("USE " + db_name)
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT INTO todo(task, dueby, status) values(?,?,?)")
	_, err = stmt.Exec(task, dueby, 1)
	if err != nil {
		panic(err)
	}
	return true

}

func (d DBUtils) GetTodoAll(db_name string) []record {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
		fmt.Println("There was a problem connecting to the db", err)
	}

	defer db.Close()

	_, err = db.Exec("USE " + db_name)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT * FROM todo")
	var recordList []record
	dateToday := time.Now()
	layout := "2006-01-02"
	for rows.Next() {
		var task string
		var dueby string
		var status bool
		var id int
		if err := rows.Scan(&id, &task, &dueby, &status); err != nil {
			fmt.Println("Error")
		}
		var record record
		t, _ := time.Parse(layout, dueby)
		completed := compareDate(t, dateToday)
		record.Dueby = dueby
		record.Task = task
		if completed {
			record.Status = "Upcoming"
		} else {
			record.Status = "Incomplete"
		}
		recordList = append(recordList, record)
	}
	if err != nil {
		panic(err)
	}
	return recordList
}

func (d DBUtils) GetTodoByDate(db_name string, date string) []record {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
		fmt.Println("There was a problem connecting to the db", err)
	}

	defer db.Close()

	_, err = db.Exec("USE " + db_name)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT * FROM todo where dueby=?", date)
	var recordList []record
	dateToday := time.Now()
	layout := "2006-01-02"
	for rows.Next() {
		var task string
		var dueby string
		var status bool
		var id int
		if err := rows.Scan(&id, &task, &dueby, &status); err != nil {
			fmt.Println("Error")
		}
		var record record
		t, _ := time.Parse(layout, dueby)
		completed := compareDate(t, dateToday)
		record.Dueby = dueby
		record.Task = task
		if completed {
			record.Status = "Upcoming"
		} else {
			record.Status = "Incomplete"
		}
		recordList = append(recordList, record)
	}
	if err != nil {
		panic(err)
	}
	return recordList
}

func compareDate(dueDate time.Time, dateToday time.Time) bool {
	if dueDate.Year() < dateToday.Year() {
		return false
	} else if dueDate.Year() > dateToday.Year() {
		return true
	} else if dueDate.Month() < dateToday.Month() {
		return false
	} else if dueDate.Month() > dateToday.Month() {
		return true
	} else if dueDate.YearDay() < dateToday.YearDay() {
		return false
	} else {
		return true
	}
}
