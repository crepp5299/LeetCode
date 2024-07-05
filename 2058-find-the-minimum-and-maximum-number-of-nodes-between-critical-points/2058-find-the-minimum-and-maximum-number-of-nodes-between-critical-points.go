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
    count := 0
    arr := []int{}
    for cur != nil && cur.Next != nil {
        if (cur.Val > prevVal && cur.Val > cur.Next.Val) ||  (cur.Val < prevVal && cur.Val < cur.Next.Val){
            arr = append(arr, count)
        }
        prevVal = cur.Val
        cur = cur.Next
        count++
    }
    l := len(arr)
    if l < 2 {
        return []int{-1,-1}
    }
    result := make([]int, 2)
    result[0] = arr[1] - arr[0]
    result[1] = 0
    for i := 1; i< l; i++ {
        if arr[i] - arr[i-1] < result[0] {
            result[0] = arr[i] - arr[i-1]
        }
    }
    result[1] = arr[l-1] - arr[0]
    
    return result
}