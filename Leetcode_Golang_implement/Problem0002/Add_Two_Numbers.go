package Problem0002

/**
题目
	You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4) Output: 7 -> 0 -> 8
 */
 
 /**
 	思路：
 		
  */
type ListNode struct {
	Val  int
	Next *ListNode
}
func Add_Two_Numbers(l1 *ListNode,l2 *ListNode) *ListNode {
	result := &ListNode{}

	temp := result

	v,n :=0,0

	for{
		// 在当前位上进行加法运算
		v, n = add(l1, l2, n)
		temp.Val = v

		// 进入下一位
		l1 = next(l1)
		l2 = next(l2)
		// 如果两个数的下一位都为nil，则结束按位相加的运算
		if l1 == nil && l2 == nil {
			break
		}

		// 为下一位运算准备节点
		temp.Next = &ListNode{}
		temp = temp.Next
	}
	// n == 1 说明，最后一次加运算进位了，需要再添加一个节点。
	if n == 1 {
		temp.Next = &ListNode{Val: n}
	}

	return result

}

func next(l *ListNode) *ListNode{
	if  l!=nil{
		return l.Next
	}
	return nil
}

func add(n1,n2 *ListNode,i int)(v,n int){
	if n1 !=nil {
		v += n1.Val
	}

	if n2 !=nil{
		v += n2.Val
	}

	v += i

	if v>9{
		v -= 10
		n = 1
	}

	return
}