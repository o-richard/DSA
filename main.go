package main

import (
	"fmt"

	"github.com/o-richard/DSA/tree"
)

func main() {
	// Three main commands
	// 1. Clean Out My Solution For A Specific File Path (Even a dir)
	// 2. Restore My Solution For A Specific File Path (Even a dir)
	// 3. Test out a package or all packages
	// Specification of design patterns
	nums := []int{7,8,9,1,2,3,4,5,7,8,8,8,10,11, -100}
	mytree := tree.InitAVL()
	for _, i := range nums {
		tree.Insert(&mytree, i)
	}
	for _, i := range nums {
		fmt.Println(tree.Search(&mytree, i))
	}
	fmt.Println(tree.Tranverse(&mytree, "preorder"))
	fmt.Println(tree.Tranverse(&mytree, "inorder"))
	fmt.Println(tree.Tranverse(&mytree, "postorder"))
	fmt.Println(tree.Tranverse(&mytree, "test"))
	fmt.Println(tree.Delete(&mytree, 2))
	fmt.Println(tree.IsBalanced(&mytree))
	// sorting an empty array, negative numbers,
	fmt.Println("Hello World")
}
