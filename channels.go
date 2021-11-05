package main

import "fmt"

/*
	permite comunicar goroutines con otras por un canal
*/
func main() {
	channel := make(chan string)

	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)

	for {
		msg := <-channel
		fmt.Println("Nombre ingresado: " + msg)
	}
}
