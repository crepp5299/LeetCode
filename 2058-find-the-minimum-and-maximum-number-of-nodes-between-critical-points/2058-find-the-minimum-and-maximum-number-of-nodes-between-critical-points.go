/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
    cur := head
    prevVal := head.Val
    cur = cur.Next
    count := 1
    first, prev, min := 0, -1, -1
    for cur != nil && cur.Next != nil {
        if (cur.Val > prevVal && cur.Val > cur.Next.Val) ||  (cur.Val < prevVal && cur.Val < cur.Next.Val){
            if min == -1 && first != 0 {
                min = count - prev
            }
            if count - prev < min {
                min = count - prev
            }
            if prev == -1 {
                first = count
            }
            prev = count
        }
        prevVal = cur.Val
        cur = cur.Next
        count++
    }
    if prev == first {
        return []int{-1,-1}
    }
    return []int{min, prev-first}

}