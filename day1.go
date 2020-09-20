package main

// TaxRate defines the tax rate for every range of income
var TaxRate = [][2]uint64{{0, 100}, {1, 100}, {2, 100}, {3, 100}}

// income > rate[k][0] && income < rate[k][1] means TaxRate[k] will be charged
// TaxRange returns range
var TaxRange = [][2]uint64{{0, 10000}, {10000, 20000}, {20000, 30000}, {30000, 40000}}

// CalculateTax return how much tax will be charged
func CalculateTax(income uint64) uint64 {
	ri := 0
	for i, tr := range TaxRange {
		if income > tr[0] && income < tr[1] {
			ri = i
		}
	}
	return income * TaxRate[ri][0] / TaxRate[ri][1]
}