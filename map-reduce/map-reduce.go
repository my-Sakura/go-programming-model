package main

import (
	"fmt"
)

// This design model is for decoupling control logic and business 
func MapStrToStr(slice []string, fn func(string) string) []string {
	var newSlice []string = make([]string, 0)
	for _, v := range slice {
		newSlice = append(newSlice, fn(v))
	}

	return newSlice
}

func Filter(slice []string, fn func(s string) bool) []string {
	var newSlice = make([]string, 0)
	for _, v := range slice {
		if fn(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func Reduce(slice []string, fn func(s string) int) int {
	sum := 0
	for _, v := range slice {
		sum += fn(v)
	}

	return sum
}

func main() {
	list := []string{"sakura", "whirlwinder", "gangfeng"}
	newList := MapStrToStr(list, func(s string) string {
		return s + "goodbye"
	})

	newListAfterReduce := Reduce(list, func(s string) int {
		return len(s)
	})

	newListAfterFliter := Filter(list, func(s string) bool {
		if len(s) > 1 {
			return true
		}
		return false
	})

	fmt.Println(newList, newListAfterFliter, newListAfterReduce)
}
