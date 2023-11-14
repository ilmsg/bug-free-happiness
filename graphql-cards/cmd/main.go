package main

import (
	"encoding/json"
	"fmt"
	"graphql-cards/cards"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	schema, err := cards.Setup()
	if err != nil {
		log.Fatal(err)
	}

	query := `
		{
			cards(value: "6"){
				value
				suit
			}
		}
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}

	rJSON, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", rJSON)
}
