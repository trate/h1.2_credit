package credit_test

import (
	"fmt"
	"github.com/trate/h1.2_credit/pkg/credit"
)

func ExampleCalculate() { // имя функции - Example + имя проверяемой функции
	fmt.Println(credit.Calculate(1_000_000_00, 36, 20))
	fmt.Println(credit.Calculate(10_000_00, 36, 20))
	// Output:
	// 3716358 33788888 133788888
	// 37163 337868 1337868
}
