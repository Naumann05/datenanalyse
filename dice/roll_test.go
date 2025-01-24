package dice

import (
	"fmt"
)

func Example_Dice() {
	x := RollSingleDieOnce()
	fmt.Println(x)
	//Output:
}

func Example_multipledice() {
	fmt.Println(RollMultipleDiceOnce(5))
	//Output:
}

func Example_manydice() {
	fmt.Println(RollMany(3, 3))
	//Output:
}
