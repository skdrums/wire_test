package main

import "fmt"

type Message string


func NewMessage() Message{
	return Message("Hi, This is TestForWireDaizin.")
}

type Greeter struct{
	Message Message
}

func NewGreeter(message Message)Greeter{
	return Greeter{Message:message}
}

func(g Greeter) Greet() Message{
	return g.Message
}

type Event struct{
	Greeter Greeter
}

func NewEvent(g Greeter) Event{
	return Event{Greeter:g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
