package main

import "fmt"

func main() {
	variable := invertirString("Holaa")
	fmt.Println(variable)
}

func invertirString(s string) string {
	new_s := ""

	j := len(s) - 1

	for i := 0; i < len(s); i++ {
		new_s += string(s[j])
		j--

	}

	return new_s
}
