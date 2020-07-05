// Writing a basic HTTP server is easy using the
// `net/http` package.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {

	type FruitBasket struct {
		Name    string
		Fruit   []string
		ID      int64  `json:"ref"`
		private string // An unexported field is not encoded.
		Created time.Time
	}

	basket := FruitBasket{
		Name:    "Standard",
		Fruit:   []string{"Apple", "Banana", "Orange"},
		ID:      999,
		private: "Second-rate",
		Created: time.Now(),
	}

	var jsonData []byte
	jsonData, err := json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

}
