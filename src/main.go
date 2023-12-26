package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArguments(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 3
}

func doubleReturn(a int) (c, d int) {
	return a, a * 3
}
func main() {
	normalFunction("hola mundo")
	tripleArguments(1, 2, "Hola")
	value := returnValue(2)
	fmt.Println("value: ", value)

	value1, value2 := doubleReturn(value)
	fmt.Println("value1 y value2: ", value1, value2)

	value3, _ := doubleReturn(value)
	fmt.Println("value3: ", value3)
}
