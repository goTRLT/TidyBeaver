package main

import "fmt"

func main() {
	var ebt, profit, ratio, taxRate, tax float64

	fmt.Print("Earnings Before Tax (EBT): ")
	fmt.Scanln(&ebt)

	for ebt == 0 {
		fmt.Println("Earnings Before Tax (EBT) cannot be zero. Please enter a valid value.")
		fmt.Scanln(&ebt)
	}

	fmt.Print("Tax Rate: ")
	fmt.Scanln(&taxRate)

	if taxRate == 0 {
		taxRate = 2.5
		fmt.Println("Tax Rate is set to default value of 2.5%")
	}

	if taxRate > 0 {
		tax = ebt * taxRate / 100
	} else {
		tax = 0
	}

	profit = ebt - tax
	if profit < 0 {
		fmt.Println("Profit did not occur since tax was bigger than earnings: ", profit)
	}

	ratio = ebt / profit
	if profit != 0 {
		ratio = ebt / profit
	}
	if ratio < 0 {
		fmt.Println("Ratio is negative since profit is negative: ", ratio)
	}

	fmt.Println("Earnings After Tax (Profit):", profit)
	fmt.Println("Ratio (EBT/Profit):", ratio)
}
