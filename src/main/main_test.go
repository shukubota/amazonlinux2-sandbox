package main_test

import (
	"github.com/shukubota/amazonlinux2-sandbox/src/main"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnqueue(t *testing.T) {
	t.Run("enqueue", func(t *testing.T) {
		err := main.Enqueue()
		assert.NoError(t, err)
	})
}
