package main_test

import (
	"github.com/shukubota/amazonlinux2-sandbox/src/main"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnqueueAndDeQueue(t *testing.T) {
	t.Run("enqueue", func(t *testing.T) {
		err := main.SendMessageToSQS("test message")
		assert.NoError(t, err)

		var message string
		message, err = main.ReceiveMessageFromSQS()
		assert.NoError(t, err)

		assert.Equal(t, message, "test message")
	})
}
