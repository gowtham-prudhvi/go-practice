package main

import "fmt"
func merge(arr []int, l int, r int, mid int) {
	return
}

func mergeSort(arr []int, l int, r int) {
	var mid int = (l + r)/2
	if (l < r) {
		fmt.Println("middle value ", arr[mid], "mid ", mid)
		mergeSort(arr, l, mid)
		mergeSort(arr, mid + 1, r)
		merge(arr, l, r, mid)
	}
	return
}

func main() {
	var a [5]int
	a[4] = 14
	fmt.Println("a[1] val:", a[1])
	fmt.Println("a[4] val:", a[4])
	fmt.Println("Hello World!!!")
	mergeSort(a[:], 0, len(a) - 1)
}