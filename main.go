package main

import (
	"fmt"
)

func multiply(a int, b int) int {
	return a * b
}

func variableArgs(str ...string) {
	defer fmt.Println("is done!")
	fmt.Println(str)
}

func main() {
	// variableArgs("gwak", "hee", "won")
	// variableArgs("gwakheewon")

	// a := 2
	// b := &a
	// a = 10
	// fmt.Println(a, *b)

	// name := [5]string{"gwak", "heewon", "won"}
	// name[3] = "천"
	// name[4] = "재"
	// fmt.Println(name)

	// slice := []string{"gwak", "heewon", "won"}
	// slice = append(slice, "천")
	// slice = append(slice, "재")
	// slice = append(slice, "다")
	// fmt.Println(slice)

	// map_test := map[string]string{"name": "heewon", "age": "23"}
	// for key, value := range map_test {
	// 	fmt.Println(key, value)
	// }

	// type person struct {
	// 	name    string
	// 	age     int
	// 	favFood []string
	// }

	// favFood := []string{"bread", "chicken"}
	// heewon := person{
	// 	name:    "heewon",
	// 	age:     23,
	// 	favFood: favFood,
	// }
	// fmt.Println(heewon)

	// account := banking.NewAccount("heewon")
	// account.Deposit(10)
	// err := account.Withdraw(11)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(account.Balance())

	// dict := mydict.Dictionary{}

	// err := dict.Add("key", "value")
	// if err == nil {
	// 	fmt.Println(dict)
	// }

	// err2 := dict.Delete("key")
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// v, err3 := dict.Search("key")

	// if err3 == nil {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println(err3)
	// }
}
