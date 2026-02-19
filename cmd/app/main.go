package main

import (
	"fmt"

	"github.com/erd2/lab4-variant12/pkg/mobiletariff"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

func main() {

	userID := uuid.New()
	fmt.Printf("User ID: %s\n\n", userID)

	minutes := 120.0
	pricePerMin := 2.5

	gb := 15.0
	pricePerGb := 450.0

	discountPercent := 10.0

	callCost, err := mobiletariff.CallCost(minutes, pricePerMin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	internetCost, err := mobiletariff.InternetCost(gb, pricePerGb)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	total := callCost + internetCost

	err = mobiletariff.ApplyPackageDiscount(&total, discountPercent)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	report, err := mobiletariff.FormatTariffReport(
		"Magzhan",
		callCost,
		internetCost,
		total,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	color.Cyan("\n===== MOBILE TARIFF REPORT =====")
	fmt.Println(report)

	fmt.Printf("\nFinal total: %10.2f ₸\n", total)
}
