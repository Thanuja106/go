package main

import "fmt"

func add(x float32, y int) float32 {
	return x + float32(y)
}
func add1(x float64, y float64) float64 {
	return x + y
}
func add3(x float64, y float32) float64 {
	return x + float64(y)
}

func add2(x float32)int {
return int(x)
}
func add31(x float64, y float32) int {
	return int(x) + int(y)
}

func main() {
	fmt.Println(add(4.33, 13))
	fmt.Println(add1(4.33, 13))
	fmt.Println(add3(4.33, 13))
	fmt.Println(add2(4.333))
	fmt.Println(add31(4.33, 13.65))
}
