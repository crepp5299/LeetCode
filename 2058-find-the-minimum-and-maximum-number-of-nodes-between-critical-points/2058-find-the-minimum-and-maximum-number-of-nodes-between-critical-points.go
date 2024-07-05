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
    first, prev, min := 0, 0, 999999 
    for cur != nil && cur.Next != nil {
        if (cur.Val > prevVal && cur.Val > cur.Next.Val) ||  (cur.Val < prevVal && cur.Val < cur.Next.Val){
            if prev == 0 {
                first = count
                prev = count
                continue
            }
            if count - prev < min && count != prev{
                min = count - prev
            }
           
            prev = count
        }
        prevVal = cur.Val
        cur = cur.Next
        count++
    }
    if prev == first || first == 0{
        return []int{-1,-1}
    }
    return []int{min, prev-first}

}