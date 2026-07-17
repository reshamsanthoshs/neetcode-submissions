/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := ListNode{}
	newlst := &dummy
    for list1!= nil && list2!= nil {
		if list1.Val < list2.Val{
			newlst.Next = list1
			list1 = list1.Next
		}else{
			newlst.Next = list2
			list2 = list2.Next
		}
		newlst = newlst.Next
	}
	if list1!= nil {
		newlst.Next = list1
	}else{
		newlst.Next = list2
	}
	return dummy.Next
}
