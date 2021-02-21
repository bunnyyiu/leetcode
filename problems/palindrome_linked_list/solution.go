/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    size := 0
    for cur := head; cur != nil; cur = cur.Next {
        size += 1
    }
    
    medium := size / 2
    
    cur := head
    for i := 0; i < medium; i ++ {
        cur = cur.Next
    }
    
    // start from medium, reverse the list
    var pre *ListNode = nil
    for cur != nil && cur.Next != nil {
        nextNext := cur.Next.Next
        next := cur.Next
        
        cur.Next.Next = cur
        cur.Next = pre
        
        pre = next
        cur = nextNext
    }
    
    if cur != nil {
        cur.Next = pre
        pre = cur
    }
    
    cur = pre
    for cur != nil {
        cur = cur.Next
    }
    
    cur1 := head
    cur2 := pre
    
    result := true
    for i := 0; i < medium; i++ {
        if cur1.Val != cur2.Val {
            result = false
        }
        cur1 = cur1.Next
        cur2 = cur2.Next
    }
    
    // fix revered list
    cur = pre
    var pre2 *ListNode = nil
    for cur != nil && cur.Next != nil {
        nextNext := cur.Next.Next
        next := cur.Next
        
        cur.Next.Next = cur
        cur.Next = pre2
        
        pre2 = next
        cur = nextNext
    }
    if cur != nil {
        cur.Next = pre2
        cur = cur.Next
    }
    
    return result
}