package hot100

/**
 * @Description
给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0开头。

 * @Author ZzzWw
 * @Date 2022-07-22 22:01
 **/

 type ListNode struct {
	    Val int
	     Next *ListNode
	 }
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
      dummyHead:=&ListNode{-1,nil}
	  temp:=dummyHead
	  carry:=0
	for l1!=nil ||l2!=nil{
		a,b:=0,0
		if l1!=nil{
			a=l1.Val
		}else{
			a=0
		}
		if l2!=nil{
			b=l2.Val
		}else{
			b=0
		}

		node:=&ListNode{Val: (a+b+carry)%10,Next: nil}
		temp.Next=node
		temp=node

		carry=(a+b+carry)/10
		if l1!=nil{
			l1=l1.Next
		}
		if l2!=nil{
			l2=l2.Next
		}
	}
	if carry>0{
		temp.Next=&ListNode{Val: carry,Next: nil}
	}
	return dummyHead.Next
}