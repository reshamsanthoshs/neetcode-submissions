/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    var nl []*ListNode
	
	if head == nil || head.Next == nil {
		return
	}

	i := 0 
	dummy := head
	for dummy!= nil {
		nl = append(nl,dummy)
		dummy = dummy.Next 
	}
	j := len(nl) - 1
	for i < j {
		next := nl[i].Next
		nl[i].Next = nl[j]
        i++

		if i > j {
			break
		}

    	nl[j].Next = next
    	j--
	}
	nl[i].Next = nil



	

}
