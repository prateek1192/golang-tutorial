package main

import "fmt"

type bot interface{
	getGreeting() string

}

type englishBot struct {
}

type spanishBot struct {
}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot){	
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string{
	// VERY CUSTOM LOGIC FOR GENERATING ENGLISH GREETING
	// Thing to note, Eb not used but compiler not complaining.
	return "Hi There!!"
}

func (sb spanishBot) getGreeting() string{
	return "Hola!"
}