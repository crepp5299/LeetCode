/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    rsTemp := head
    sum := 0
    temp := head.Next
    for temp != nil {
        if temp.Val == 0 {
            temp.Val = sum
            sum = 0
            rsTemp.Next = temp
            rsTemp = temp
        } else {
            sum+= temp.Val
        }
        temp = temp.Next
        
    }
    return head.Next
}