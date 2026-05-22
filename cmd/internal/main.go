package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

type body struct {
	otherData map[string]string
}

type app struct {
	logger  *slog.Logger
	url     string
	body    body
	method  string
	workers int
	length  int
}

func main() {

	otp := make(chan string, 100) // make it buffered so it is decoupled from the workers

	var wg sync.WaitGroup // assigning it's to it's type zero value

	var app app
	flag.StringVar(&app.url, "url", "none", "The url")
	flag.StringVar(&app.method, "method", "POST", "The method to be using")
	flag.IntVar(&app.workers, "t", 10, "The threads to be used")
	flag.IntVar(&app.length, "l", 4, "The length of the otp")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app.logger = logger

	fmt.Printf("Working with %d many threads", app.workers)

	for i := 0; i < app.workers; i++ {
		wg.Add(1)
		go app.callRequest(i, otp, &wg)
	}

	go app.generateOtp(otp)

	wg.Wait()
}
