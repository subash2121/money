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

func (money *Money) AddPaise(paise int) {
	if money.GetPaise()+paise%100 >= 100 {
		money.rupee += (money.GetPaise() + paise) / 100
		money.paise = (money.GetPaise() + paise) % 100
	} else if paise >= 100 {
		money.rupee += paise / 100
		money.paise += paise % 100
	} else {
		money.paise += paise
	}
}

func (money *Money) AddMoney(moneyTwo Money) {
	money.AddRupee(moneyTwo.GetRupee())
	money.AddPaise(moneyTwo.GetPaise())
}
