package main

type Money struct {
	rupee int64
	paise int64
}

func NewMoney() Money {
	return Money{}
}

func (money Money) GetRupee() int64 {
	return money.rupee
}

func (money Money) GetPaise() int64 {
	return money.paise
}
