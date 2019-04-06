package main

// 队列一个一个合并， 扩展了2个队列合并的代码。

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, end *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	t1, t2 := l1, l2

	if l1.Val < l2.Val {
		head, end = l1, l1
		t1 = t1.Next
	} else {
		head, end = l2, l2
		t2 = t2.Next
	}
	end.Next = nil

	for t1 != nil && t2 != nil {
		if t1.Val < t2.Val {
			end.Next = t1
			end = end.Next
			t1 = t1.Next
			end.Next = nil
		} else {
			end.Next = t2
			end = end.Next
			t2 = t2.Next
			end.Next = nil
		}
	}

	if t1 != nil {
		end.Next = t1
	}
	if t2 != nil {
		end.Next = t2
	}

	return head
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	var head *ListNode
	for _, l := range lists {
		head = mergeTwoLists(l, head)
	}
	return head
}

func main() {

}
