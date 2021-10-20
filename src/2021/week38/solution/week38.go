package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type todo struct {
	Id          int
	UserId      int
	Title       string
	IsCompleted bool `json:"completed"`
}

func main() {
	responseObj, err := makeHttpRequest("https://jsonplaceholder.typicode.com/todos/4", &todo{})
	if err != nil {
		fmt.Printf("failed to get todo: %v", err)
		return
	}

	todo := responseObj.(*todo)

	fmt.Printf("Todo: %+v\n", todo)
}

func makeHttpRequest(url string, responseTemplate interface{}) (interface{}, error) {
	var body []byte

	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to read body: %v", err))
		}
	}

	// // NOTE: Go playground users will have to uncomment this section and replace the "INSERT_JSON_OBJECT_HERE"
	// // placeholder with the actual JSON object being returned by the endpoint used in the main function (i.e. visit the link).
	// // This is due to Go playground not allowing external requests.
	// jsonStr := `INSERT_JSON_OBJECT_HERE`
	// body = []byte(jsonStr)

	err = json.Unmarshal(body, &responseTemplate)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to unmarshal json: %v", err))
	}

	return responseTemplate, nil
}
