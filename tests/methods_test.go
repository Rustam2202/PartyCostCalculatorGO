package tests

import (
	"party-calc/internal"
	"party-calc/utils"
	"testing"
)

func TestShowPayments(t *testing.T) {
	utils.IntializeLogger()
	utils.LoadConfig()
	result := internal.CalculateDebts(case1.Input, 1)
	result.ShowPayments()
}
