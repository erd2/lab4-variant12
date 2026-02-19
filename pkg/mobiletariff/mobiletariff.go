import "github.com/erd2/lab4-variant12/pkg/mobiletariff"

// Package mobiletariff предоставляет функции для расчёта стоимости
// мобильного тарифа: звонки, интернет, скидки и формирование отчёта.
package mobiletariff

import (
	"fmt"
)

// CallCost вычисляет стоимость звонков.
//
// minutes — количество минут.
// pricePerMin — цена за минуту.
//
// Возвращает общую стоимость или ошибку при некорректных данных.
func CallCost(minutes float64, pricePerMin float64) (float64, error) {
	if minutes < 0 {
		return 0, fmt.Errorf("minutes cannot be negative")
	}
	if pricePerMin < 0 {
		return 0, fmt.Errorf("price per minute cannot be negative")
	}

	return minutes * pricePerMin, nil
}

// InternetCost вычисляет стоимость интернет-трафика.
//
// gb — объём интернета в гигабайтах.
// pricePerGb — цена за 1 ГБ.
//
// Возвращает стоимость или ошибку.
func InternetCost(gb float64, pricePerGb float64) (float64, error) {
	if gb < 0 {
		return 0, fmt.Errorf("gigabytes cannot be negative")
	}
	if pricePerGb < 0 {
		return 0, fmt.Errorf("price per GB cannot be negative")
	}

	return gb * pricePerGb, nil
}

// ApplyPackageDiscount применяет скидку к общей сумме.
//
// total — указатель на итоговую стоимость.
// percent — процент скидки.
//
// Изменяет значение по указателю.
func ApplyPackageDiscount(total *float64, percent float64) error {
	if total == nil {
		return fmt.Errorf("total pointer is nil")
	}
	if percent < 0 || percent > 100 {
		return fmt.Errorf("discount percent must be between 0 and 100")
	}

	discount := *total * percent / 100
	*total -= discount

	return nil
}

// FormatTariffReport формирует строку отчёта по тарифу.
//
// user — имя клиента.
// calls — стоимость звонков.
// internet — стоимость интернета.
// total — итоговая сумма.
//
// Возвращает форматированный отчёт.
func FormatTariffReport(user string, calls, internet, total float64) (string, error) {
	if user == "" {
		return "", fmt.Errorf("user name cannot be empty")
	}

	report := fmt.Sprintf(
		"Tariff report for %s\nCalls: %.2f ₸\nInternet: %.2f ₸\nTotal: %.2f ₸",
		user,
		calls,
		internet,
		total,
	)

	return report, nil
}
