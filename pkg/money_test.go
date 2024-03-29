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

func TestInputValidation(t *testing.T) {
	t.Run("should panic if negative inputs are given", func(t *testing.T) {
		assert.Panics(t, func() {
			InputValidation(-1)
		})
	})

	t.Run("should return true if positive inputs are given", func(t *testing.T) {
		assert.Equal(t, true, InputValidation(1))
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
	t.Run("should add 10 and 5 rupee to the existing money", func(t *testing.T) {
		money := NewMoney()
		money.AddRupee(10)
		assert.Equal(t, 10, money.GetRupee())
		money.AddRupee(5)
		assert.Equal(t, 15, money.GetRupee())
	})

	t.Run("should panic if -5 rupee is added", func(t *testing.T) {
		money := NewMoney()
		assert.Panics(t, func() {
			money.AddRupee(-5)
		})
	})
}

func TestSubtractRupee(t *testing.T) {
	t.Run("should subtract 10 and 5 rupee to the existing money of 20 rupee", func(t *testing.T) {
		money := NewMoney()
		money.AddRupee(20)
		money.SubtractRupee(10)
		assert.Equal(t, 10, money.GetRupee())
		money.SubtractRupee(5)
		assert.Equal(t, 5, money.GetRupee())
	})

	t.Run("should panic if -5 rupee is subtracted", func(t *testing.T) {
		money := NewMoney()
		assert.Panics(t, func() {
			money.SubtractRupee(-5)
		})
	})
}

func TestAddPaise(t *testing.T) {
	t.Run("should add 50 paise to the new money", func(t *testing.T) {
		money := NewMoney()
		money.AddPaise(50)
		assert.Equal(t, 50, money.GetPaise())
	})

	t.Run("should add 120 paise to the new money as 1 rupee and 20 paise", func(t *testing.T) {
		money := NewMoney()
		money.AddPaise(120)
		assert.Equal(t, 1, money.GetRupee())
		assert.Equal(t, 20, money.GetPaise())
	})

	t.Run("should add 170 paise to existing money having 50 paise as 2 rupee and 20 paise", func(t *testing.T) {
		money := NewMoney()
		money.AddPaise(50)
		money.AddPaise(170)
		assert.Equal(t, 2, money.GetRupee())
		assert.Equal(t, 20, money.GetPaise())
	})

	t.Run("should add 400 paise to existing money having 2 rupee 50 paise as 6 rupee and 50 paise", func(t *testing.T) {
		money := NewMoney()
		money.AddRupee(2)
		money.AddPaise(50)
		money.AddPaise(400)
		assert.Equal(t, 6, money.GetRupee())
		assert.Equal(t, 50, money.GetPaise())
	})

	t.Run("should panic if -50 paise is added", func(t *testing.T) {
		money := NewMoney()
		assert.Panics(t, func() {
			money.AddPaise(-50)
		})
	})
}

func TestSubtractPaise(t *testing.T) {
	t.Run("should subtract 50 paise to the new money", func(t *testing.T) {
		money := NewMoney()
		money.SubtractPaise(50)
		assert.Equal(t, -50, money.GetPaise())
	})

	t.Run("should subtract 120 paise to the new money as -1 rupee and -20 paise", func(t *testing.T) {
		money := NewMoney()
		money.SubtractPaise(120)
		assert.Equal(t, -1, money.GetRupee())
		assert.Equal(t, -20, money.GetPaise())
	})

	t.Run("should subtract 170 paise to existing money having 50 paise as -1 rupee and -20 paise", func(t *testing.T) {
		money := NewMoney()
		money.AddPaise(50)
		money.SubtractPaise(170)
		assert.Equal(t, -1, money.GetRupee())
		assert.Equal(t, -20, money.GetPaise())
	})

	t.Run("should subtract 400 paise to existing money having 2 rupee 50 paise as -1 rupee and -50 paise", func(t *testing.T) {
		money := NewMoney()
		money.AddRupee(2)
		money.AddPaise(50)
		money.SubtractPaise(400)
		assert.Equal(t, -1, money.GetRupee())
		assert.Equal(t, -50, money.GetPaise())
	})

	t.Run("should panic if -70 paise is subtracted", func(t *testing.T) {
		money := NewMoney()
		assert.Panics(t, func() {
			money.SubtractPaise(-70)
		})
	})
}

func TestAddMoney(t *testing.T) {
	t.Run("should add money with 1 rupee and 50 paise with 2 rupee and 70 paise", func(t *testing.T) {
		moneyOne := NewMoney()
		moneyOne.AddRupee(1)
		moneyOne.AddPaise(50)

		moneyTwo := NewMoney()
		moneyTwo.AddRupee(2)
		moneyTwo.AddPaise(70)

		moneyOne.AddMoney(moneyTwo)
		assert.Equal(t, 4, moneyOne.GetRupee())
		assert.Equal(t, 20, moneyOne.GetPaise())
	})
}

func TestSubtractMoney(t *testing.T) {
	t.Run("should subtract money with 1 rupee and 50 paise with 1 rupee and 30 paise", func(t *testing.T) {
		moneyOne := NewMoney()
		moneyOne.AddRupee(1)
		moneyOne.AddPaise(50)

		moneyTwo := NewMoney()
		moneyTwo.AddRupee(1)
		moneyTwo.AddPaise(30)

		moneyOne.SubtractMoney(moneyTwo)
		assert.Equal(t, 0, moneyOne.GetRupee())
		assert.Equal(t, 20, moneyOne.GetPaise())
	})
}

func TestMoneyEqualsTo(t *testing.T) {
	t.Run("should return true for money 50 paise and money 50 paise", func(t *testing.T) {
		moneyOne := NewMoney()
		moneyOne.AddPaise(50)

		moneyTwo := NewMoney()
		moneyTwo.AddPaise(50)

		assert.Equal(t, true, moneyOne.EqualsTo(moneyTwo))
		assert.Equal(t, true, moneyTwo.EqualsTo(moneyOne))
	})

	t.Run("should return true for money 1 rupee 60 paise and money 160 paise", func(t *testing.T) {
		moneyOne := NewMoney()
		moneyOne.AddRupee(1)
		moneyOne.AddPaise(60)

		moneyTwo := NewMoney()
		moneyTwo.AddPaise(160)

		assert.Equal(t, true, moneyOne.EqualsTo(moneyTwo))
		assert.Equal(t, true, moneyTwo.EqualsTo(moneyOne))
	})

	t.Run("should return true for money 170 paise and money 1 rupee 70 paise", func(t *testing.T) {
		moneyOne := NewMoney()
		moneyOne.AddPaise(170)

		moneyTwo := NewMoney()
		moneyTwo.AddRupee(1)
		moneyTwo.AddPaise(70)

		assert.Equal(t, true, moneyOne.EqualsTo(moneyTwo))
		assert.Equal(t, true, moneyTwo.EqualsTo(moneyOne))
	})

	t.Run("should return true for money 1 rupee 50 paise and money 1 rupee 50 paise", func(t *testing.T) {
		moneyOne := NewMoney()
		moneyOne.AddRupee(1)
		moneyOne.AddPaise(50)

		moneyTwo := NewMoney()
		moneyTwo.AddRupee(1)
		moneyTwo.AddPaise(50)

		assert.Equal(t, true, moneyOne.EqualsTo(moneyTwo))
		assert.Equal(t, true, moneyTwo.EqualsTo(moneyOne))
	})

	t.Run("should return false for money 1 rupee 50 paise and money 1 rupee 60 paise", func(t *testing.T) {
		moneyOne := NewMoney()
		moneyOne.AddRupee(1)
		moneyOne.AddPaise(50)

		moneyTwo := NewMoney()
		moneyTwo.AddRupee(1)
		moneyTwo.AddPaise(60)

		assert.Equal(t, false, moneyOne.EqualsTo(moneyTwo))
		assert.Equal(t, false, moneyTwo.EqualsTo(moneyOne))
	})

	t.Run("should return true for money 2 rupee,50 paise and money 250 paise and money 1 rupee,150 paise", func(t *testing.T) {
		moneyOne := NewMoney()
		moneyOne.AddRupee(2)
		moneyOne.AddPaise(50)

		moneyTwo := NewMoney()
		moneyTwo.AddPaise(250)

		moneyThree := NewMoney()
		moneyThree.AddRupee(1)
		moneyThree.AddPaise(150)

		assert.Equal(t, true, moneyOne.EqualsTo(moneyTwo))
		assert.Equal(t, true, moneyTwo.EqualsTo(moneyThree))
		assert.Equal(t, true, moneyThree.EqualsTo(moneyOne))
	})
}
