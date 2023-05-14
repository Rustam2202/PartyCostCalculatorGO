package tests

import (
	"party-calc/internal"
	"party-calc/utils"
	"testing"
)

func TestShowPayments(t *testing.T) {
	utils.IntializeLogger()
	utils.LoadConfig()
	result := internal.CalculateDebts(threePersons.input, 1)
	result.ShowPayments()
}
