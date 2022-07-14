package main

type Money struct {
	rupee int64
	paise int64
}

func NewMoney() Money {
	return Money{}
}
