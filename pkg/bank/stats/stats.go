package stats

import (
	"github.com/Buzurgmehr92/bank/pkg/bank/types"
)

//Avg рассчитывает средюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	i := types.Money(0)
	for _, payment := range payments {
		i++
		sum = sum + payment.Amount
	}
	return sum / i
}

//TotalInCategory находит сумму покупок в определённой категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category {
			sum = sum + payment.Amount
		}
	}
	return sum
}
