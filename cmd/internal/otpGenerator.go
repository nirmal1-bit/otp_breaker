package main

import "fmt"

func (app *app) generateOtp(otp chan string) {

	var maxValue int = 1
	for i := 0; i < app.length; i++ {
		maxValue *= 10
	}
	maxValue -= 1

	for i := 0; i < maxValue; i++ {
		otp <- fmt.Sprintf("%0*d", app.length, i)
	}
	close(otp)
}
