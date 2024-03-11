package main

import "fmt"

type companyProfit struct {
	ID         string
	Profit     int
	Department string
}

type DepartMentProfit struct {
	Departement string
	Profit      int
}

type CompanyTotalProfit struct {
	CompanyID   string
	TotalProfit int
	Department  []DepartMentProfit
}

func CalculateTotalProfit(cmp_pr []companyProfit) []CompanyTotalProfit {

	CompTotalProfits := make(map[string]CompanyTotalProfit) // creating map

	for _, cmp_profit := range cmp_pr {

		if cmp, ok := CompTotalProfits[cmp_profit.ID]; ok {
			cmp.TotalProfit += cmp_profit.Profit

			index := -1

			for i, dep := range cmp.Department {

				if dep.Departement == cmp_profit.Department {
					index = i
					break
				}

			}

			if index != -1 {

				cmp.Department[index].Profit += cmp_profit.Profit

			} else {
				depProfit := DepartMentProfit{
					Departement: cmp_profit.Department,
					Profit:      cmp_profit.Profit,
				}

				cmp.Department = append(cmp.Department, depProfit)

			}
			CompTotalProfits[cmp_profit.ID] = cmp

		} else {

			new_comp := CompanyTotalProfit{
				CompanyID:   cmp_profit.ID,
				TotalProfit: cmp_profit.Profit,
				Department:  []DepartMentProfit{{Departement: cmp_profit.Department, Profit: cmp_profit.Profit}},
			}

			CompTotalProfits[cmp_profit.ID] = new_comp

		}

	}

	var res []CompanyTotalProfit

	for _, company := range CompTotalProfits {
		res = append(res, company)
	}

	return res

}

func main() {

	companyProfits := []companyProfit{
		{ID: "1", Profit: 10, Department: "A"},
		{ID: "1", Profit: 5, Department: "A"},
		{ID: "1", Profit: 5, Department: "B"},
		{ID: "2", Profit: 5, Department: "A"},
	}

	// Calculate total profits
	totalProfits := CalculateTotalProfit(companyProfits)

	// Print the result
	fmt.Printf("%+v\n", totalProfits)

}
