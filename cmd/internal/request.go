package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (app *app) makeRequest(otp string) {

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	jsonData := fmt.Appendf(nil, `{"token":"%s","newPassword":"somerandomepass","confirmNewPassword":"somerandomepass"}`, otp)
	req, err := http.NewRequest(app.method, app.url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", "hostname")
	if err != nil {
		app.logger.Error(err.Error())
	}

	res, err := client.Do(req)

	if err != nil {
		app.logger.Error(err.Error())
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

}
