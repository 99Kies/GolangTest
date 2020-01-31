package main

import (
	"fmt"
	"reflect"
)

func killRepetion(nums [][]int) [][]int  {
	newRes := make([][]int, 0)
    for i := 0; i < len(nums); i++ {
        flag := false
        for j := i + 1; j < len(nums); j++ {
           if reflect.DeepEqual(nums[i], nums[j]){
                flag = true
                break
            }
        }
        if !flag {
            newRes = append(newRes, nums[i])
        }
    }
	return newRes
}

func main()  {
	result := [][]int{{1,2,3},{1,2,3},{1,2,3},{1,2,2}}
	killDoble := killRepetion(result)
	fmt.Println(killDoble)
}