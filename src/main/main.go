package main

import (
	"fmt"
	adaptor "github.com/shukubota/amazonlinux2-sandbox/src/sqs_adaptor"
)

func main() {
	err := Main()
	fmt.Println(err)
}

func Main() error {
	err := SendMessageToSQS("test")
	if err != nil {
		return err
	}
	_, err = ReceiveMessageFromSQS()
	if err != nil {
		return err
	}
	return nil
}

const queueName = "prepared"

func SendMessageToSQS(message string) error {
	s, err := adaptor.NewAdapter()
	if err != nil {
		return err
	}

	err = s.Enqueue(queueName, message)
	if err != nil {
		return err
	}
	return nil
}

func ReceiveMessageFromSQS() (string, error) {
	s, err := adaptor.NewAdapter()
	if err != nil {
		return "", err
	}

	messages, err := s.Dequeue(queueName)
	if err != nil {
		return "", err
	}

	if len(messages) == 0 {
		return "", nil
	}

	message := *messages[0].Body
	return message, nil
}
