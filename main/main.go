package main

import "homework"
import "fmt"

func main() {
	a := [15]float32{1,2,3,4,5,6,0}
	fmt.Println(homework.Average(a))
	b := []int64{1,2,5,15}
	fmt.Println(homework.Reverse(b))
	//c := map[int]string{2: "a", 0: "b", 1: "c"}
	c := map[int]string{2: "a", 0: "b", 1: "c"}
	fmt.Println(homework.SortMapValues(c))
}