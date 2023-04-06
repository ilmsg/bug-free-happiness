package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	apiUrl := "https://jsonplaceholder.typicode.com/users"
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Name", "Email"})
	for _, user := range users {
		writer.Write([]string{
			fmt.Sprintf("%d", user.ID),
			fmt.Sprintf("%s", user.Name),
			fmt.Sprintf("%s", user.Email),
		})
	}
}
