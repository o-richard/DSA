package main

import (
	"fmt"
	"github.com/o-richard/DSA/linkedlist"
)

func main() {
	// Three main commands
	// 1. Clean Out My Solution For A Specific File Path (Even a dir)
	// 2. Restore My Solution For A Specific File Path (Even a dir)
	// 3. Test out a package or all packages
	mylist := linkedlist.InitCDLL()
	linkedlist.Insertion(&mylist, 4, 0)
	linkedlist.Insertion(&mylist, 2, 0)
	linkedlist.Insertion(&mylist, 3, 2)
	linkedlist.Insertion(&mylist, 7, 2)
	fmt.Println(linkedlist.Tranverse(&mylist, false))
	fmt.Println(linkedlist.Tranverse(&mylist, true))
	index, found := linkedlist.Retrieval(&mylist, 2, false)
	fmt.Println(linkedlist.Tranverse(&mylist, true))
	fmt.Println(linkedlist.Tranverse(&mylist, false))
	fmt.Println(index, found)
	fmt.Println("Hello World")
}
