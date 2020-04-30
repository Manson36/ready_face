package main

import (
	"fmt"
)

func bubble(li []int) {
	for i := 0; i < len(li) - 1; i++ {
		for j := 0; j < len(li) - 1 - i; j++ {
			if li[j] > li[j+1] {
				li[j], li[j+1] = li[j+1], li[j]
			}
		}
	}
}

func select_sort(li []int) {
	for i := 0; i < len(li) -1; i++ {
		for j := i+1; j < len(li); j++ {
			if li[i] > li[j] {
				li[i], li[j]= li[j], li[i]
			}
		}
	}
}

func insert_sort(li []int) {
	for i := 1; i < len(li); i++ {
		tmp := li[i]
		j := i - 1
		for j >= 0 && tmp < li[j] {
			li[j+1] = li[j]
			j--
		}
		li[j + 1] = tmp
	}
}

func insert_sort2(li []int) {
	for i:=1; i<len(li); i++ {
		tmp := li[i]
		j := i-1
		for j >= 0 && tmp < li[j] {
			li[j+1] = li[j]
			j--
		}
		li[j + 1] = tmp
	}
}

func quick(li []int) {
	if len(li) <= 1 {
		return
	}

	left, right := 0, len(li) - 1
	mid, i := li[left], 1

	for left < right {
		if mid > li[i] {
			li[i], li[left] = li[left], li[i]
			i++
			left++
		} else {
				li[i], li[right] = li[right], li[i]
				right--
		}
	}

	fmt.Println(li)
	quick(li[:i])
	quick(li[i:])
}

func main() {
	ss := []int{22, 3, 12, 1, 55, 9}

	//bubble(ss)
	//select_sort(ss)
	//insert_sort(ss)
	//insert_sort2(ss)
	quick(ss)
	fmt.Println(ss)
}
