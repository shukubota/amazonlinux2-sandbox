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

func (s *Adapter) SendMessageToQueue() {
	fmt.Println("sendMessage")
}

func (s *Adapter) Dequeue() {
	fmt.Println("aaa")
}

func (s *Adapter) GetList() error {
	fmt.Println("getlist")
	input := &sqs.GetQueueAttributesInput{
		AttributeNames: []*string{aws.String("All")},
		QueueUrl:       aws.String("prepared"),
	}
	m, err := s.sqs.GetQueueAttributes(input)
	fmt.Println("---------------")
	fmt.Println(m)
	fmt.Println(err)
	fmt.Println("---------------")
	if err != nil {
		return err
	}
	fmt.Println(m)
	return nil
}
