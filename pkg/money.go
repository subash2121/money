package main

const unitConversion = 100

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

func (money *Money) SubtractRupee(rupee int) {
	money.rupee -= rupee
}

func (money *Money) AddPaise(paise int) {
	if money.GetPaise()+paise >= unitConversion {
		money.AddRupee((money.GetPaise() + paise) / unitConversion)
		money.paise = (money.GetPaise() + paise) % unitConversion
	} else {
		money.paise += paise
	}
}

func (money *Money) SubtractPaise(paise int) {
	if money.GetPaise()-paise <= -unitConversion {
		money.AddRupee((money.GetPaise() - paise) / unitConversion)
		money.paise = (money.GetPaise() - paise) % unitConversion
	} else {
		money.paise -= paise
	}
}

func (money *Money) AddMoney(moneyTwo Money) {
	money.AddRupee(moneyTwo.GetRupee())
	money.AddPaise(moneyTwo.GetPaise())
}

func (money Money) EqualsTo(moneyTwo Money) bool {
	return money == moneyTwo
}
