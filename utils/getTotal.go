package utils

import "pwaa-test.com/models/entity"

func GetTotal(items []entity.Pwaa) int {
	var totalWorth = 0

	for _, pwaa := range items {
		totalWorth += pwaa.Worth
	}
	
	return totalWorth
}