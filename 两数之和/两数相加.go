package main

import (
	"fmt"

)

func twoSum(nums []int, target int) []int {
	n := len(nums)
	result := []int{}
    for i:=0; i<n-1; i++{
        for j:=i+1; j<n; j++{
			if nums[i]+nums[j] == target{
				result = append(result, i, j)
				return result
			}
		}
	}
	return result
}

func twoSum_hash(nums []int, target int) []int {
	tables := make(map[int]int) // 声明一个空字典

	for index, value := range nums {
		// index为地址，value为nums[index]的值
		tmp := target - value
		fmt.Printf("index %d\tvalue %d\ttmp %d\n",index, value, tmp)
		i, ok := tables[tmp]
		fmt.Printf("tmp %d\ti %d\tok %v\n", tmp,i,ok)
		if i, ok := tables[tmp]; ok {
			return []int{i, index}
		}
		tables[value] = index
	}
	return []int{}
}

func main() {
	var nums = []int{3,2,4}
	done := twoSum(nums, 6)
	fmt.Println(done)
	// var nums_hash = []int{2,7,11,15}
	// done_hash := twoSum_hash(nums_hash, 9)
	// fmt.Println(done_hash)
}



package main

import "fmt"


func threeSum(nums []int) [][]int {
	var temp int = 0
	for n:=0; n<len(nums)-1; n++{
		for m:=n+1; m<len(nums); m++{
			if nums[m]<nums[n]{
				temp = nums[m]
				nums[m] = nums[n]
				nums[n] = temp
			}
		}
	}
	fmt.Println(nums)
	result := [][]int{}
    for i:=0; i<len(nums)-2; i++{
		res := []int{}
		for j:=i+1; j<len(nums)-1; j++{
			for k:=j+1; k<len(nums); k++{
				if nums[i]+nums[j]+nums[k] == 0{
					// fmt.Println(i,j,k)
					res  = append(res, nums[i], nums[j], nums[k])
					// break
				}
			}
			if len(res)>0 {
				fmt.Println(res)
				result = append(result,  []int{res[0], res[1], res[2]})
				fmt.Println(res)
				fmt.Println(result)
			}
			res = res[0:0]
			
		}
	}
	fmt.Println(result)
	newArr := make([][]int, 0)
    for a := 0; a < len(result); a++ {
        repeat := false
        for b := a + 1; b < len(result); b++ {
            if result[a][0] == result[b][0] && result[a][1] == result[b][1] && result[a][2] == result[b][2]{
                repeat = true
                break
            }
        }
        if !repeat {
            newArr = append(newArr, result[a])
        }
    }
	return result
}

func main()  {
	nums := []int{-1,0,1,2,-1,-4}

	done := threeSum(nums)
	
	fmt.Println(done)
	
}