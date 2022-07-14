package main

type Money struct {
	rupee int
	paise int
}

func NewMoney() Money {
	return Money{}
}

func (money Money) GetRupee() int {
	return money.rupee
}

func (money Money) GetPaise() int {
	return money.paise
}

func (money *Money) AddRupee(rupee int) {
	money.rupee += rupee
}
