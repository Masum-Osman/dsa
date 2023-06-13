package tree

import "fmt"

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

/* LC 104 */
func BinaryTreeMaxDepth(root *treeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := BinaryTreeMaxDepth(root.Left)
	rightHeight := BinaryTreeMaxDepth(root.Right)

	leftHeight++
	rightHeight++

	if leftHeight > rightHeight {
		fmt.Println("left")
		return leftHeight
	} else {
		fmt.Println("right")
		return rightHeight
	}

}

/* LC 231 */
func IsPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	if n%2 != 0 || n <= 0 {
		return false
	}

	return IsPowerOfTwo(n / 2)

}

func IsPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n%3 != 0 || n <= 0 {
		return false
	} else {
		return IsPowerOfThree(n / 3)
	}
}

func IsPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}
	if n%4 != 0 || n <= 0 {
		return false
	} else {
		return IsPowerOfFour(n / 4)
	}
}

/* LC 344 */
func ReverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/* LC 700 */
func SearchBST(root *treeNode, val int) *treeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	if root.Val > val {
		return SearchBST(root.Left, val)
	} else {
		return SearchBST(root.Right, val)
	}
}

/* LC 938 */
func RangeSumBST(root *treeNode, low int, high int) int {
	sum := 0

	return helper(root, low, high, &sum)
}

func helper(root *treeNode, low int, high int, sum *int) int {
	if root == nil {
		return 0
	}
	if root.Val >= low && root.Val <= high {
		*sum = *sum + root.Val
	}
	helper(root.Left, low, high, sum)
	helper(root.Right, low, high, sum)

	return *sum
}

func RangeSumBST0(root *treeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	fmt.Println(low, high, root.Val)
	sum := 0
	if root.Val >= low && root.Val <= high {
		fmt.Println("Inside IF. Val: ", root.Val, "Sum: ", sum)
		sum = sum + root.Val
	}
	RangeSumBST(root.Left, low, high)
	RangeSumBST(root.Right, low, high)

	return sum
}

/* LC 2236 */
func getNodeValue(root *treeNode) int {
	if root == nil {
		return 0
	}
	return root.Val
}

func IfEqual(root *treeNode) bool {
	if root == nil {
		return false
	}
	return root.Val == getNodeValue(root.Left)+getNodeValue(root.Right)
}

func CheckTree(root *treeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
