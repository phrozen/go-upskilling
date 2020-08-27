package main

import (
	"container/list"
	"fmt"
)

type tree struct {
	value int
	left  *tree
	right *tree
}

func createNode(i int) *tree {
	var root tree
	root.value = i
	return &root
}
func createBTS(n int) *tree {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i + 1
	}
	return createBTSHelper(a)
}
func createBTSHelper(s []int) *tree {
	l := len(s)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return createNode(s[0])
	}
	i := l / 2
	root := createNode(s[i])
	root.left = createBTSHelper(s[:i])
	root.right = createBTSHelper(s[i+1:])

	return root
}

func createTree(n int) *tree {
	if n == 0 {
		return nil
	}
	l := list.New()
	root := createNode(1)
	l.PushBack(root)
	for i := 2; i < n+1; i++ {
		tmpRoot := l.Front()
		l.Remove(tmpRoot)
		tmp := createNode(i)
		tmpRoot.Value.(*tree).left = tmp
		if i == n {
			return root
		}
		l.PushBack(tmp)
		i++
		tmp = createNode(i)
		tmpRoot.Value.(*tree).right = tmp
		l.PushBack(tmp)
	}
	return root
}

func preorderTransversalRecursion(root *tree) {
	if root == nil {
		return
	}
	fmt.Println(root.value)
	preorderTransversalRecursion(root.left)
	preorderTransversalRecursion(root.right)
}

func preorderTransversalNoRecursion(root *tree) {
	//We create a new list to store the nodes to visit
	l := list.New()
	//Since the list is waiting a empty interface we can pass our struct without problems
	l.PushBack(root)
	for l.Len() > 0 {
		//This will return an *Element
		tmp := l.Front()
		//We do a little of reflection and cast back from Element
		tmpNode := tmp.Value.(*tree)
		l.Remove(tmp)
		fmt.Println(tmpNode.value)
		if tmpNode.right != nil {
			l.PushFront(tmpNode.right)
		}
		if tmpNode.left != nil {
			l.PushFront(tmpNode.left)
		}
	}
}

func canSum(a []int, t int) bool {
	m := make(map[int]bool)
	for _, i := range a {
		if _, ok := m[t-i]; ok {
			return true
		}
		m[i] = true
	}
	return false
}

func main() {
	root := createTree(3)
	preorderTransversalRecursion(root)
	preorderTransversalNoRecursion(root)
	root = createBTS(3)
	fmt.Println("BTS")
	preorderTransversalRecursion(root)
	fmt.Println("BTS nonrecursive")
	preorderTransversalNoRecursion(root)
	a := make([]int, 30)
	for i := 0; i < 30; i++ {
		a[i] = i + 1
	}
	fmt.Println(canSum(a, 3))
}
