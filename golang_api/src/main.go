package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
)

type Item struct {
	Title string `json:"title"`
	Url string `json:"url"`
}

func main() {
	response, err := http.Get("https://qiita.com/api/v2/items")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var items []Item
	marshalErr := json.Unmarshal(responseData, &items)
	if marshalErr != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Printf("title:%v, url:%v\n", item.Title, item.Url)
	}
}