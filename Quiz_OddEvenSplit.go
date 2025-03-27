package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList1(head *ListNode) *ListNode {
	//两种节点分别为奇数节点和偶数节点，额外空间为o(1)仅仅表示不能有额外的节点，但是原有节点的指向改变是没有问题的
	dummyHeadOdd := &ListNode{}
	dummyHeadEven := &ListNode{}
	currOdd := dummyHeadOdd
	currEven := dummyHeadEven
	flagOddEven := 1 //1 = odd, 0 = even
	for head != nil {
		if flagOddEven == 1 {
			currOdd.Next = head
			currOdd = currOdd.Next
		} else {
			currEven.Next = head
			currEven = currEven.Next
		}
		head = head.Next
		flagOddEven = flagOddEven ^ 1
	}
	currOdd.Next = dummyHeadEven.Next
	currEven.Next = nil
	return dummyHeadOdd.Next
}

func oddEvenList2(head *ListNode) *ListNode {
	dummyHeadOdd := &ListNode{}
	dummyHeadEven := &ListNode{}
	currOdd := dummyHeadOdd
	currEven := dummyHeadEven
	for head != nil {
		currOdd.Next = head
		currOdd = currOdd.Next
		head = head.Next
		if head != nil {
			currEven.Next = head
			currEven = currEven.Next
		} else {
			break
		}
		head = head.Next
	}
	currOdd.Next = dummyHeadEven.Next
	currEven.Next = nil
	return dummyHeadOdd.Next
}

func main() {
	//testLinkNode
	lk1 := &ListNode{Val: 1}
	lk2 := &ListNode{Val: 5}
	lk3 := &ListNode{Val: 2}
	lk4 := &ListNode{Val: 6}
	lk5 := &ListNode{Val: 3}
	lk6 := &ListNode{Val: 7}
	lk7 := &ListNode{Val: 4}
	// lk8 := &ListNode{Val: 8}
	lk1.Next = lk2
	lk2.Next = lk3
	lk3.Next = lk4
	lk4.Next = lk5
	lk5.Next = lk6
	lk6.Next = lk7
	lk7.Next = nil
	// lk7.Next = lk8
	// lk8.Next = nil
	lk1 = oddEvenList1(lk1)
	for lk1 != nil {
		println(lk1.Val)
		lk1 = lk1.Next
	}

}
