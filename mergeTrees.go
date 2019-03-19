package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    var t3 *TreeNode = new(TreeNode)
    if t1 == nil {
        if (t2 != nil) {
            return t2
        } else {
            return nil
        }
    } else {
        if (t2 == nil) {
            return t1
        }
        t3.Val = t1.Val + t2.Val
        t3.Left =mergeTrees(t1.Left, t2.Left)
        t3.Right =mergeTrees(t1.Right, t2.Right)
        return t3
    }
}

func addNode(val int) *TreeNode {
    var t *TreeNode = new(TreeNode)
    t.Val = val
    return t
}

func t1() *TreeNode {
    var t1 *TreeNode = new(TreeNode)
    t1.Val = 1
    t1.Left = addNode(3)
    t1.Right = addNode(2)
    t1.Left.Left = addNode(5)
    return t1
}

func t2() *TreeNode {
    var t2 *TreeNode = new (TreeNode)
    t2.Val = 2
    t2.Left = addNode(1)
    t2.Right = addNode(3)
    t2.Left.Right = addNode(4)
    t2.Right.Right = addNode(7)
    return t2
}

func printTree(t *TreeNode) {
    if t == nil {
        return
    }
    fmt.Print(t.Val, "-")
    printTree(t.Left)
    printTree(t.Right)
    fmt.Println()
}

func main() {
   t1 := t1()
   t2 := t2()
   printTree(t1) 
   printTree(t2) 
   printTree(mergeTrees(t1, t2))
}


