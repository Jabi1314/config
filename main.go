package main

import (
	"fmt"

	"github.com/aspirations/aspirations-mapper/mapper"
)

func main() {
	// First Task
	str := "Aspiration.com"
	fmt.Println("Before mapper..", str)
	capStr := mapper.CapitalizeEveryThirdAlphanumericChar(str)
	fmt.Println("After Mapper..", capStr)

	// Second Task
	s := mapper.NewSkipString(3, "Aspiration.com")
	fmt.Println("Before mapper..", s)
	mapper.MapString(&s)
	fmt.Println("After Mapper..", s)
}
