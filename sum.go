package sum

import "fmt"

func DoSum() {

	var i int
	var j int

	fmt.Print("Type your numbers: ")
	fmt.Scan(&i, &j)
	sumofNum := i + j
	fmt.Printf("Sum of the two numbers is %v\n", sumofNum)

}
