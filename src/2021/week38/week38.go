/**
 * Code challenge week 38, 2021 (language: GO, playground: https://play.golang.org/)
 *
 * Description:
 *     The code below is supposed to fetch a todo object from an external HTTP endpoint and then print it. So we run the code and, well,
 *     the object is printed nice and smoothly. But since we are good engineers we also pay a visit to the source endpoint (by opening the
 *     URL in our favourite browser). And when we do this we realise something terrifying. The source object doesn’t match what our
 *     application is printing! Dear lord - dark times ahead…
 *
 * Questions:
 *     1. Why is there a mismatch between the source object and what is printed?
 *     2. How can it be adjusted to print a match of the source object?
 */

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
	IsCompleted bool
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
	// // placeholder with the actual JSON object being returned by the endpoint used in the main function (i.e. visit the url).
	// // This is due to Go playground not allowing external requests.
	// jsonStr := `INSERT_JSON_OBJECT_HERE`
	// body = []byte(jsonStr)

	err = json.Unmarshal(body, &responseTemplate)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to unmarshal json: %v", err))
	}

	return responseTemplate, nil
}
