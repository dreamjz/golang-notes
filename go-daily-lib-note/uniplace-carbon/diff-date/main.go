package main

import (
	"fmt"
	"github.com/uniplaces/carbon"
)

func main() {
	outputFormat := "%-30s:%t\n"

	date1, _ := carbon.CreateFromDate(2010, 1, 1, "Asia/Shanghai")
	date2, _ := carbon.CreateFromDate(2011, 2, 1, "Asia/Shanghai")
	date3, _ := carbon.CreateFromDate(2010, 12, 1, "Asia/Shanghai")
	fmt.Println("date1:", date1)
	fmt.Println("date1:", date2)
	fmt.Println("date1:", date3)
	fmt.Printf(outputFormat, "date1 equal to  date2", date1.Eq(date2))
	fmt.Printf(outputFormat, "date1 not equal to date2", date1.Ne(date2))

	fmt.Printf(outputFormat, "date1 greater than date2", date1.Gt(date2))
	fmt.Printf(outputFormat, "date1 less than date2", date1.Lt(date2))

	fmt.Printf(outputFormat, "date3 between date1 and date2", date3.Between(date1, date2, true))

	now := carbon.Now()
	fmt.Printf("%-30s:%s\n", "now", now)
	fmt.Printf(outputFormat, "is weekday", now.IsWeekday())
	fmt.Printf(outputFormat, "is weekend", now.IsWeekend())
	fmt.Printf(outputFormat, "is leap year", now.IsLeapYear())
	fmt.Printf(outputFormat, "is past", now.IsPast())
	fmt.Printf(outputFormat, "is future", now.IsFuture())
}
