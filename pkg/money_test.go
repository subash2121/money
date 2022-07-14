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

func TestGetRupee(t *testing.T) {
	t.Run("should return the rupees", func(t *testing.T) {
		assert.Equal(t, int64(0), NewMoney().GetRupee())
	})
}

func TestGetPaise(t *testing.T) {
	t.Run("should return the paise", func(t *testing.T) {
		assert.Equal(t, int64(0), NewMoney().GetPaise())
	})
}
