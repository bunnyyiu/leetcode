/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }
    result := l1
    
    
    if l1.Val < l2.Val {
        l1 = l1.Next
    } else {
        result = l2
        l2 = l2.Next
    }
    
    curResult := result
    
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            curResult.Next = l1
            l1 = l1.Next
        } else {
            curResult.Next = l2
            l2 = l2.Next
        }
        curResult = curResult.Next
    }
    
    for l1 != nil {
        curResult.Next = l1
        curResult = curResult.Next

        l1 = l1.Next
    }
    
    for l2 != nil {
        curResult.Next = l2
        curResult = curResult.Next

        l2 = l2.Next
    }
    
    return result
}