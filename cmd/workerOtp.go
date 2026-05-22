package main

import (
	"fmt"
	"sync"
)

func (app *app) callRequest(id int, otp chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	for num := range otp {
		fmt.Printf("Worked id %d, otp value %s \n", id, num)
		app.makeRequest(num)
	}

}
