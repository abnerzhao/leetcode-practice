/*
链表中环的检测 https://leetcode-cn.com/problems/linked-list-cycle/

快慢指针，有环一定会相遇

 Definition for singly-linked list.
 type ListNode struct {
      Val int
      Next *ListNode
 }
 */

 func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    slow, fast := head, head.Next
    for slow != fast {
        if fast == nil || fast.Next == nil {
            return false
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    return true
}