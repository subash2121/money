package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMoney(t *testing.T) {
	t.Run("should be able to initialize and return money entity", func(t *testing.T) {
		assert.IsType(t, Money{}, NewMoney())
	})
}
