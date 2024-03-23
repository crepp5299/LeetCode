/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    m := map[int]int{}
    count := 0
    for head != nil {
        m[count] = head.Val
        head = head.Next
        count++
    }
    i:=0
    for i < count/2 {
        if m[i] != m[count-i-1]{
            return false
        }
        i++
    }
    return true
}