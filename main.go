package main

import (
	"fmt"
	"github.com/o-richard/DSA/heap"
)

func main() {
	// Three main commands
	// 1. Clean Out My Solution For A Specific File Path (Even a dir)
	// 2. Restore My Solution For A Specific File Path (Even a dir)
	// 3. Test out a package or all packages
	nums := []int{1, 2, 3, 4, 5,6,7}
	nums1 := []int{1, 2, 3, 4, 5,6,7}
	myheap := heap.InitFHeap(false)
	myheap2 := heap.InitFHeap(true)
	for _, i := range nums {
		myheap.InsertFHeap(i)
	}
	for _, i := range nums1 {
		myheap2.InsertFHeap(i)
	}
	nice, _ := myheap.ExtractRootFHeap()
	val, _ := myheap2.ExtractRootFHeap()
	fmt.Println(nice, val)
	fmt.Println()
	myheap.UnionFHeap(&myheap2)
	fmt.Println("Hello World")
	for _, i := range nums {
		new := myheap.Test(i)
		booli := new != nil	
		fmt.Println(booli)
	}
	for _, i := range nums1 {
		new := myheap.Test(i)
		booli := new != nil	
		fmt.Println(booli)
	}
	fmt.Println(myheap.DeleteNode(4))
	fmt.Println(myheap.Test(4))
	fmt.Println(myheap.DeleteNode(6))
	myheap.ChangeKey(5, 1)
	fmt.Println(myheap.Test(5))
	nice, _ = myheap.FindRootFHeap()
	val, _ = myheap.FindRootFHeap()
	fmt.Println()
	fmt.Println(nice, val)
}
