package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func (app *app) makeRequest(otp string) {

	client := http.Client{}

	jsonData := fmt.Sprintf(`{
    "reset_token":"%s",
    "password":"whatisthis",
    "conform":"whatisthis"
}`, otp)

	req, err := http.NewRequest(app.method, app.url, strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		app.logger.Error(err.Error())
		log.Fatal("Error server did not response", err)
	}

	res, err := client.Do(req)

	if err != nil {
		app.logger.Error(err.Error())
		log.Fatal("Error server did not response", err)
	}

	defer res.Body.Close()
	// io.ReadCloser interface  which has  Reader and Closer method defined

	// io.ReadAll()
	// and this takes Reader interface and the res.Body satisfies this as res.Body is ReadCloser interface which is defined like this
	//
	//type ReadCloser interface {
	// Reader
	// Closer
	// }

	data, err := io.ReadAll(res.Body)

	if err != nil {
		app.logger.Error(err.Error())

		return
	}

	fmt.Printf("Status: %d \n", res.StatusCode)
	fmt.Println(string(data))

	if res.StatusCode == 200 {

		fmt.Printf("\nThe otp is %s", otp)
		log.Fatal()
	}

}
