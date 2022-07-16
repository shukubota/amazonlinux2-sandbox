package main

import (
	"fmt"
	"github.com/shukubota/amazonlinux2-sandbox/src/sqs"
)

func main() {
	Enqueue()
}

func Enqueue() {
	fmt.Println("testしたい")
	s, err := sqs.NewAdapter()
	if err != nil {
		fmt.Println(err)
		return
	}

	s.SendMessageToQueue()
}
