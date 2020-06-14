// Package credit provides annuity payments calculation
package credit

import (
	"math"
	"fmt"
)

func Calculate(creditSum, termInMonths, ratePerYear uint64) (monthlyPayment, overpayment, allPayments uint64) {

	ratePerMonth := float64(ratePerYear) / 12.0 / 100.0
	k := ratePerMonth * math.Pow(1 + ratePerMonth, float64(termInMonths)) / (math.Pow(1 + ratePerMonth, float64(termInMonths)) - 1)
	fmt.Println(k)
	monthlyPayment = uint64(float64(creditSum) * k)
	allPayments = monthlyPayment * termInMonths
	overpayment = allPayments - creditSum

	return
}
