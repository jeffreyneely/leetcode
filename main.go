package main

import "fmt"

func main() {
	fmt.Printf("%d\n", 23%10)
	fmt.Printf("%d\n", 321%10)
	fmt.Printf("%d\n", (321/10)%10)
	arr := []int{4, 2, 13, 4, -1, 8}
	mv := MaxValue(arr)
	fmt.Printf("maxVal: %d\n", mv)
	fmt.Printf("countElem: %d\n", CountElem(arr))
	fmt.Printf("sumElem: %d\n", SumElem(arr))
	a2 := []int{3, 21, 1, 4, 3, -1}
	QSort(&a2, 0, len(a2)-1)
	fmt.Printf("qSort: %v", a2)
}
