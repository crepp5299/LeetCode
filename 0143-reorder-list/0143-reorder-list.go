/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    m := map[int]*ListNode{}
    count := 0
    curr := head
    for curr != nil {
        m[count] = curr
        curr = curr.Next
        count++
    }
    if count <= 2 {
        return
    }
    head.Next = m[count-1]
    i := 1
    for i <=count/2 {
        if i == count/2 {
            m[i].Next = nil
            if count%2 > 0 {
               m[count-i].Next = m[i]
            }
            break
        }
        m[count-i].Next = m[i]
        m[i].Next = m[count-i-1]
        i++
    }
}