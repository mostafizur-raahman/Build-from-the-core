package clouser

import "fmt"

func activateGiftCard() func(int) int {
	amount := 100

	activateCard := func(data int) int {
		amount -= data
		return amount
	}
	return activateCard
}
func GiftSomething() {
	persone1 := activateGiftCard()
	persone2 := activateGiftCard()
	fmt.Println(persone1(10), persone2(100))
}
