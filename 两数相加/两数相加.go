package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	temp := l1
// 	for ; temp.Next != null; temp=temp.Next{
// 		fmt.Println(temp.Val)
// 	}
// }

func main() {
	var a *ListNode
	// var b *ListNode
	a.Val = 2
	a.Next.Val = 5
	a.Next.Next.Val = 8
	// for ; a.Next!=nil; a = a.Next{
	// 	fmt.Println(a.Val)
	// }
	if a.Next != nil {
		for {
			fmt.Println("\t%v", a.Val)
			if a.Next != nil {
				a = a.Next
			} else {
				break
			}
		}
	}
}
