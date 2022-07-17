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
	region := "us-east-1"
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String("localhost.localstack.cloud:4566"),
	})
	if err != nil {
		return nil, err
	}

	adapter := sqs.New(sess)
	fmt.Println(adapter)

	//fmt.Println(sqs_service)
	return &Adapter{
		sqs: adapter,
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
