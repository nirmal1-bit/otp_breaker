package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ResetRequest struct {
	ResetToken string `json:"reset_token"`
	Password   string `json:"password"`
	Conform    string `json:"conform"`
}

func (app *app) makeRequest(otp string) {
	reqBody := ResetRequest{
		ResetToken: otp,
		Password:   "batmanstuff",
		Conform:    "batmanstuff",
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("failed to marshal json: %v", err)
	}

	req, err := http.NewRequest(app.method, app.url, bytes.NewReader(jsonData))

	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		app.logger.Error(err.Error())
		log.Print("Error server did not response", err)
		return
	}

	res, err := app.client.Do(req)
	if err != nil {
		app.logger.Error(err.Error())
		log.Print("Error server did not response", err)
		return
	}

	if res.StatusCode == app.stopHttpCode {
		fmt.Printf("\nThe otp is %s", otp)
		os.Exit(0)
	}

}
