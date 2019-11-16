package main

import "fmt"

func main(){

	nums := []int{}
	for i:=1;i<=10;i++{
		nums = append(nums, i)
	}
	for _, val := range(nums){
		if val % 2 !=0 {
			fmt.Printf("%d is odd\n", val)
		} else {
			fmt.Printf("%d is even\n", val)
		}
	}
}