package adaptor

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Adapter struct {
	sqs *sqs.SQS
}

func NewAdapter() (*Adapter, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("ap-northeast-1"),
		Endpoint: aws.String("http://localhost:4566"),
	})
	if err != nil {
		return nil, err
	}
	return &Adapter{
		sqs: sqs.New(sess),
	}, nil
}

func (s *Adapter) Enqueue(queueName string, message string) error {
	i := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueName),
		MessageBody: aws.String(message),
	}
	_, err := s.sqs.SendMessage(i)
	if err != nil {
		return err
	}
	return nil
}

func (s *Adapter) Dequeue(queueName string) ([]*sqs.Message, error) {
	input := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(queueName),
	}
	o, err := s.sqs.ReceiveMessage(input)
	if err != nil {
		return nil, err
	}
	return o.Messages, nil
}

func (s *Adapter) GetQueue(queueName string) error {
	input := &sqs.GetQueueAttributesInput{
		AttributeNames: []*string{aws.String("All")},
		QueueUrl:       aws.String(queueName),
	}
	m, err := s.sqs.GetQueueAttributes(input)
	if err != nil {
		return err
	}
	fmt.Println(m)
	return nil
}
