// 删除链表节点
// https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteNode(head *ListNode, val int) *ListNode {

	if (head.Val == val){
		return head.Next
	}
	pre := head
	for head.Next.Val != val {
		head = head.Next
	}
	head.Next = head.Next.Next
	return pre
}




