// Package credit provides annuity payments calculation
package credit

import (
	"math"
)

func Calculate(creditSum, termInMonths, ratePerYear int) (int, int, int) {

	ratePerMonth := float64(ratePerYear) / 12.0 / 100.0
	// three levels of the precision
	ratePerMonth = math.Round(ratePerMonth*1_000) / 1_000
	k := ratePerMonth * math.Pow(1+ratePerMonth, float64(termInMonths)) / (math.Pow(1+ratePerMonth, float64(termInMonths)) - 1)
	// six levels of the precision
	k = math.Round(k*1_000_000) / 1_000_000
	monthlyPayment := float64(creditSum) * k
	allPayments := monthlyPayment * float64(termInMonths)
	overpayment := allPayments - float64(creditSum)

	return int(monthlyPayment), int(overpayment), int(allPayments)
}
