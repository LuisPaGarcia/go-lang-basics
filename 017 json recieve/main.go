package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	Number int `json:"number"`
}

// Bird structure
type Bird struct {
	Species     string
	Description string
}

// ToDO struct
type ToDO struct {
	UserID    int32  `json:"userId"`
	ID        int32  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	url := "https://jsonplaceholder.typicode.com/todos/"

	httpClient := http.Client{
		Timeout: time.Second * 21, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var todos []ToDO
	json.Unmarshal([]byte(body), &todos)

	for i := range todos {
		if todos[i].Completed == true {
			fmt.Println(todos[i])
		}
	}
}
