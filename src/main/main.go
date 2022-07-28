package main

import (
	adaptor "github.com/shukubota/amazonlinux2-sandbox/src/sqs_adaptor"
	"log"
)

func main() {
	log.Fatal(Main())
}

func Main() error {
	s, err := adaptor.NewAdapter()
	if err != nil {
		return err
	}
	message := "生殺与奪の権を他人に握らせるな!!"
	err = SendMessageToSQS(s, message)
	if err != nil {
		return err
	}
	_, err = ReceiveMessageFromSQS(s)
	if err != nil {
		return err
	}
	return nil
}

const queueName = "prepared"

func SendMessageToSQS(s *adaptor.Adapter, message string) error {
	err := s.Enqueue(queueName, message)
	if err != nil {
		return err
	}
	return nil
}

func ReceiveMessageFromSQS(s *adaptor.Adapter) (string, error) {
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
