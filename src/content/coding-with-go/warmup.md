---
title: Warming up
tag: coding
index: 1
---

## Warming up

Once that we have the basics concepts of the language we can start coding, to warm up we are going to solve some basic problems.

### Tree preorder transversal

The first thing that we need to do is create the structure for the tree.

``` go
type tree struct{
    value num
    left *tree
    right *tree
}
```

Once that we the struct for the tree, we can implement the algorithm for it, for the sake of this exercise we will do it first preorder recursive and then preorder without recursion.

#### Recursive

Creating Recursive solution is pretty simple since the recursion in go is the same as in other languages:

``` go
func preorderTransversalRecursion(root *tree) {
    if root == nil {
        return
    }
    fmt.Println(root.value)
    preorderTransversalRecursion(root.left)
    preorderTransversalRecursion(root.right)
}
```

#### Non recursive

To create a non recursive solution we need a queue or a list were to store the nodes to visit.

``` go
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
```

### Two numbers in array can sum a number

This problem can be solved with a map so, lets use one.

``` go
func canSum(a []int, t int) bool {
	m := make(map[int]bool)
  // we can check all elements in the array usnig a range
    for _, i := range a {
        if _, ok := m[t-i]; ok {
            return true
        }
        m[i] = true
    }
    return false
}
```
