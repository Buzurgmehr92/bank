package stats

import (
	"fmt"

	"github.com/Buzurgmehr92/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   55_00,
			Category: "Cat",
		},
		{
			ID:       1,
			Amount:   58_00,
			Category: "Cat",
		},
		{
			ID:       3,
			Amount:   30_00,
			Category: "Cat",
		},
	}

	fmt.Println(Avg(payments))

	//Output: 4766
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   63_00,
			Category: "Cafe",
		},
		{
			ID:       1,
			Amount:   23_00,
			Category: "Refuelling",
		},
		{
			ID:       3,
			Amount:   42_00,
			Category: "Refuelling",
		},
	}

	fmt.Println(TotalInCategory(payments, "Refuelling"))

	//Output: 6500
}
