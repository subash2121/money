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
		assert.Equal(t, 0, NewMoney().GetRupee())
	})
}

func TestGetPaise(t *testing.T) {
	t.Run("should return the paise", func(t *testing.T) {
		assert.Equal(t, 0, NewMoney().GetPaise())
	})
}

func TestAddRupee(t *testing.T) {
	t.Run("should add rupee to the existing money", func(t *testing.T) {
		money := NewMoney()
		money.AddRupee(10)
		assert.Equal(t, 10, money.GetRupee())
		money.AddRupee(5)
		assert.Equal(t, 15, money.GetRupee())
	})
}
