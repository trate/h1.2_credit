package main

import (
	"fmt"
	"github.com/trate/h1.2_credit/pkg/credit"
)

func main() {
	monthlyPayment, overpayment, allPayments := credit.Calculate(1_000_000_00, 36, 20)
	fmt.Println(monthlyPayment)
	fmt.Println(overpayment)
	fmt.Println(allPayments)
}
