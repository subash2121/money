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

func InputValidation(input int) bool {
	if input < 0 {
		panic("Negative input should not be provided")
	}
	return true
}

func (money *Money) AddRupee(rupee int) {
	if InputValidation(rupee) {
		money.rupee += rupee
	}
}

func (money *Money) SubtractRupee(rupee int) {
	if InputValidation(rupee) {
		money.rupee -= rupee
	}
}

func (money *Money) AddPaise(paise int) {
	if InputValidation(paise) {
		if money.GetPaise()+paise >= unitConversion {
			money.AddRupee((money.GetPaise() + paise) / unitConversion)
			money.paise = (money.GetPaise() + paise) % unitConversion
		} else {
			money.paise += paise
		}
	}
}

func (money *Money) SubtractPaise(paise int) {
	if InputValidation(paise) {
		if money.GetPaise()-paise <= -unitConversion {
			money.SubtractRupee(-(money.GetPaise() - paise) / unitConversion)
			money.paise = (money.GetPaise() - paise) % unitConversion
		} else {
			money.paise -= paise
		}
	}
}

func (money *Money) AddMoney(moneyTwo Money) {
	money.AddRupee(moneyTwo.GetRupee())
	money.AddPaise(moneyTwo.GetPaise())
}

func (money *Money) SubtractMoney(moneyTwo Money) {
	money.SubtractRupee(moneyTwo.GetRupee())
	money.SubtractPaise(moneyTwo.GetPaise())
}

func (money Money) EqualsTo(moneyTwo Money) bool {
	return money.rupee == moneyTwo.rupee && money.paise == moneyTwo.paise
}
