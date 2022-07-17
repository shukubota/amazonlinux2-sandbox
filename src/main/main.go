package main

import (
	"fmt"
	adaptor "github.com/shukubota/amazonlinux2-sandbox/src/sqs_adaptor"
)

func main() {
	Enqueue()
}

func Enqueue() error {
	s, err := adaptor.NewAdapter()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = s.GetList()
	if err != nil {
		return err
	}
	return nil
}
