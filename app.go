package main

import "fmt"

type CashFlow struct {
	netIncome            float64
	costOfGoodsSold      float64 // def negative
	opex                 float64 // def negative
	gainsFromSale        float64 // either
	depreciation         float64 // def negative
	appreciation         float64 // def positive
	interest             float64 // def negative
	ammortization        float64 // def negative
	cTaxRate             float64 // positive decimal
	changeWorkingCapital float64
}

func (cf CashFlow) EBITDA() float64 {
	return (cf.netIncome + cf.costOfGoodsSold + cf.opex + cf.gainsFromSale)
}

func (cf CashFlow) EBIT() float64 {
	return (cf.EBITDA() + cf.depreciation + cf.appreciation)
}

func (cf CashFlow) Tax() float64 {
	return (-1 * (cf.EBIT() + cf.interest) * cf.cTaxRate)
}

func (cf CashFlow) TaxShield() float64 {
	return (cf.interest * cf.cTaxRate)
}

func (cf CashFlow) EAT() float64 {
	EBT := (cf.EBIT() + cf.interest)
	return (EBT + cf.Tax())
}

func main() {

	testCf := CashFlow{
		netIncome:            1000000,
		costOfGoodsSold:      -600000,
		opex:                 -100000,
		gainsFromSale:        0,
		depreciation:         0,
		appreciation:         0,
		interest:             -150000,
		ammortization:        0,
		cTaxRate:             0.25,
		changeWorkingCapital: 0,
	}

	fmt.Println("EBIT " + fmt.Sprintf("%f", testCf.EBIT()) + " EAT " + fmt.Sprintf("%f", testCf.EAT()) + " Tax shield " + fmt.Sprintf("%f", testCf.TaxShield()))
}
