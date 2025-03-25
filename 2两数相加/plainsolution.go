/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1 := getLen(l1)
	len2 := getLen(l2)
	// p1 and p2 are moving pointers.
	var p1 *ListNode
	var p2 *ListNode
	var res *ListNode
	// if length of l1 greater than l2 , return l1, otherwise l2.
	if len1 >= len2 {
		p1 = l1
		p2 = l2
		res = l1
	} else {
		p1 = l2
		p2 = l1
		res = l2
	}
	sum := 0
	digit := 0
	// p1_stop is the previous pointer of p1.
	// when p1 runs to the end of l1, p1_stop is the last node of l1.
	var p1_stop *ListNode
	if len1 == len2 {
		for p1 != nil && p2 != nil {
			sum = p1.Val + p2.Val + digit
			if sum > 9 {
				digit = 1
			} else {
				digit = 0
			}
			p1.Val = sum % 10
			p2 = p2.Next
			if p2 == nil {
				p1_stop = p1
			}
			p1 = p1.Next
		}
		// to handle the carry condition
		if digit == 1 {
			tmp := new(ListNode)
			tmp.Val = 1
			tmp.Next = nil
			p1_stop.Next = tmp
		}
	} else {
		for p1 != nil && p2 != nil {
			sum = p1.Val + p2.Val + digit
			if sum > 9 {
				digit = 1
			} else {
				digit = 0
			}
			p1.Val = sum % 10
			p2 = p2.Next
			if p2 == nil {
				p1_stop = p1
			}
			p1 = p1.Next
		}
		if digit == 1 {
			for p1 != nil {
				sum = p1.Val + digit
				if sum > 9 {
					digit = 1
				} else {
					digit = 0
				}
				p1.Val = sum % 10
				if p1.Next == nil {
					p1_stop = p1
				}
				p1 = p1.Next
			}
			if digit == 1 {
				tmp := new(ListNode)
				tmp.Val = 1
				tmp.Next = nil
				p1_stop.Next = tmp
			}
		}
	}
	return res
}

func getLen(list *ListNode) int {
	res := 0
	for list != nil {
		res += 1
		list = list.Next
	}
	return res
}