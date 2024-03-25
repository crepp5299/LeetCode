/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    curr := list1
    b = b - a + 1
    for a > 1 {
        curr = curr.Next
        a--
    }
    nextCurr := curr
    for b > 0 {
        nextCurr = nextCurr.Next
        b--
    }
    curr.Next = list2
    for curr.Next != nil {
        curr = curr.Next
    }
    curr.Next = nextCurr.Next
    return list1
}