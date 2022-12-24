package main

import "fmt"

type engBot struct {
}

type spaBot struct {
}

type bot interface{
	getGreeting() string
}

func main() {
	eb := engBot{}
	sp := spaBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (engBot) getGreeting() string {
	return "hello"
}
func (spaBot) getGreeting() string {
	return "hola"
}
