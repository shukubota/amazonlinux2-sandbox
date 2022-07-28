package main_test

import (
	"github.com/shukubota/amazonlinux2-sandbox/src/main"
	adaptor "github.com/shukubota/amazonlinux2-sandbox/src/sqs_adaptor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnqueueAndDeQueue(t *testing.T) {
	t.Run("enqueueしたメッセージをdequeueできメッセージを受け取れる", func(t *testing.T) {
		s, err := adaptor.NewAdapter()
		assert.NoError(t, err)

		// sqsにenqueue
		err = main.SendMessageToSQS(s, "生殺与奪の権を他人に握らせるな!!")
		assert.NoError(t, err)

		var message string
		// sqsからdequeue
		message, err = main.ReceiveMessageFromSQS(s)
		assert.NoError(t, err)

		assert.Equal(t, message, "生殺与奪の権を他人に握らせるな!!")
	})
}
