package main

import "fmt"

func main() {
	minAmount := 5000
	purchase := 15000
	limitCashback := 2000
	percentCashback := 15
	cashback := 0
	const fullPercent = 100

	if purchase >= minAmount {
		cashback = purchase * percentCashback / fullPercent
		fmt.Println("Prediction cashback is", cashback)
	}

	if limitCashback > cashback {
		cashback = limitCashback
		fmt.Println("cashback is", cashback)
	}
	//----HOME WORK----
	// cashback
	// Если сумма покупки состовляет более 500смн скидка в размере 15%
	// Если сумма кешбека превышает 2000 смн, то кешбек должен быть 2000смн
}
