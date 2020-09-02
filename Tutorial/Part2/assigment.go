package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	fibNums := []int{0, 1, 2, 3, 5, 8}
	//fibNums := getFibNumbers(len(nums))
	var result map[string][]int = make(map[string][]int, 0)
	//fmt.Println(nums[5:])
	for _, v := range nums {

		if v%2 == 0 {
			if result["even"] == nil { // {"odd":[1]}
				result["even"] = make([]int, 0) // {"odd":[1],"even":[]}
			}
			result["even"] = append(result["even"], v) // {"even":[2,4,6,8,10]}
		} else {
			if result["odd"] == nil { // {}
				result["odd"] = make([]int, 0) // {"odd":[]}
			}
			result["odd"] = append(result["odd"], v) // {"odd":[1,3,4,7,9]}
		}

		for _, f := range fibNums {
			if f == v {
				if result["fib"] == nil { //{}
					result["fib"] = make([]int, 0) //{"fib":[]}
				}
				result["fib"] = append(result["fib"], v) //{"fib":[1]}
			}
		}
	}

	fmt.Println("Even :", result["even"])
	fmt.Println("Odd :", result["odd"])
	fmt.Println("Fib :", result["fib"])
}

func getFibNumbers(n int) []int {
	start := 0
	next := 1
	var fibNums []int = make([]int, 0)
	fibNums = append(fibNums, start)
	for {
		sum := start + next
		fibNums = append(fibNums, sum)
		if sum > n {
			return fibNums
		}
		start = next
		next = sum
	}
}
