package sqs

import "fmt"

type Adapter struct {
	url string
}

func NewAdapter() (*Adapter, error) {
	return &Adapter{}, nil
}

func (s *Adapter) SendMessageToQueue() {
	fmt.Println("sendMessage")
}
