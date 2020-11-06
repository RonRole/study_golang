package main

import (
	"fmt"
	"strconv"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func randomFromDB() float64 {
	db, err := sql.Open("mysql", "root:root@tcp(mysql)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var value float64;

	results, err := db.Query("SELECT RAND()")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		results.Scan(&value)
	}
	defer results.Close()

	return value;
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for i := range []int{1,2,3,4,5} {
			fmt.Fprintf(w,strconv.FormatFloat(randomFromDB(), 'f', -1, 64))
		}
	})
	http.ListenAndServe(":9090", nil)
}