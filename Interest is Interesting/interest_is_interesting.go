package interest

 InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	interestRate = 0.0
    if(balance  0){
        interestRate = 3.213
    }else if(balance = 0 && balance  1000){
    	interestRate = 0.5
    }else if(balance = 1000 && balance  5000){
    	interestRate = 1.621
    }else{
    	interestRate = 2.475
    }
	return float32(interestRate)
}

 Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return float64(InterestRate(balance))  balance100
}

 AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

 YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	yearsCount = 0
	for balance  targetBalance{
        balance = AnnualBalanceUpdate(balance)
        yearsCount++
    }
	return yearsCount
}
