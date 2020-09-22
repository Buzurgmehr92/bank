package stats

import (
  "github.com/Buzurgmehr92/bank/pkg/bank/types"
  "fmt"
)


func ExampleAvg(){
  payments := []types.Payment{
    {
      ID:2,
      Amount:55_00,
      Category: "Cat",
    },
    {
      ID:1,
      Amount:58_00,
      Category: "Cat",
    },
    {
      ID:3,
      Amount:30_00,
      Category: "Cat",
    },
  }

  fmt.Println(Avg(payments))

  //Output: 4766
}