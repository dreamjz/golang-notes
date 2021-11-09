package main

import (
	"fmt"

	"github.com/uniplaces/carbon"
)

func main() {
	date1, _ := carbon.CreateFromDate(2021, 1, 1, "Asia/Tokyo")
	date2, _ := carbon.CreateFromDate(2022, 1, 1, "Asia/Tokyo")

	fmt.Println(date1.DiffInYears(date2, false))   // 1
	fmt.Println(date1.DiffInMonths(date2, false))  // 12
	fmt.Println(date1.DiffInDays(date2, false))    // 365
	fmt.Println(date1.DiffInHours(date2, false))   // 8760
	fmt.Println(date1.DiffInMinutes(date2, false)) // 525600
	fmt.Println(date1.DiffInSeconds(date2, false)) // 3153600

}
