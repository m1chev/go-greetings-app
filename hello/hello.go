package main

import (
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	switch len(os.Args[1:]) {
	case 0:
		fmt.Println("Please provide an argument!")
	case 1:
		message, err := greetings.Hello(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(message)
	default:
		messages, err := greetings.Hellos(os.Args[1:])
		if err != nil {
			log.Fatal(err)
		}
		for _, gr := range messages {
			fmt.Println(gr)
		}
	}
}

/*
	if len(os.Args[1:]) > 0 {
		message, err := greetings.Hello(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(message)
	} else {
		fmt.Println("Please provide an argument!")
		fmt.Println(len(os.Args[1:]))
	}
*/
