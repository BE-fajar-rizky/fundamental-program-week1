package main

import "fmt"

func DrawXYZ(n int) string {
	// your code here
	var pangkat int = 2
	var bobot string
	var testt int = 1
	for i := 0; i < pangkat; i++ {
		testt *= n
	}
	for i := 1; i <= testt; i++ {

		if i%3 == 0 {
			bobot += "X "

		} else if i%2 == 0 {
			bobot += "Z "
		} else {
			bobot += "Y "
		}
		if i%n == 0 {
			bobot += "\n"
		}
	}
	return bobot
}
func main() {
	fmt.Println(DrawXYZ(3))
	fmt.Println(DrawXYZ(5))
}
