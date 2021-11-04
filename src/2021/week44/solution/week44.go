package main

import (
	"errors"
	"fmt"
)

type municipality struct {
	ID                int
	TotalYearlyIncome int
	Population        int
}

func main() {
	municipalities := []municipality{
		{
			ID:                1,
			TotalYearlyIncome: 30000000,
			Population:        100,
		},
		{
			ID:                2,
			TotalYearlyIncome: 0,
			Population:        0,
		},
	}

	for _, municipality := range municipalities {
		yearlyAverageIncome, err := getYearlyAverageMunicipalityIncome(municipality)

		if err != nil {
			rootErr := unwrapErr(err)

			switch rootErr.Error() {
			case "Can't divide by zero":
				fallbackAverageIncome := 0
				yearlyAverageIncome = &fallbackAverageIncome
			default:
				fmt.Printf("A general error occurred: %s\n", rootErr)
				return
			}
		}

		fmt.Printf("The yearly average income for municipality %d is %d\n", municipality.ID, *yearlyAverageIncome)
	}
}

func getYearlyAverageMunicipalityIncome(municipality municipality) (*int, error) {
	averageIncome, err := divide(municipality.TotalYearlyIncome, municipality.Population)

	if err != nil {
		return nil, fmt.Errorf("Could not calculate yearly average municipality income: %w", err)
	}

	return &averageIncome, nil
}

func divide(firstNumber int, secondNumber int) (int, error) {
	if secondNumber == 0 {
		return 0, errors.New("Can't divide by zero")
	}

	return firstNumber / secondNumber, nil
}

func unwrapErr(err error) error {
	u := errors.Unwrap(err)
	if u == nil {
		return err
	}
	return unwrapErr(u)
}
