// 链表中倒数第K个节点
// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // getKthFromEnd 所有链表节点放入列表，取倒数第K个
 // 此方法空间复杂度为O(n)
func getKthFromEnd(head *ListNode, k int) *ListNode {
	var s []*ListNode
    for head != nil {
		s = append(s, head)
		head = head.Next
	}
	if len(s) >= k {
		return s[len(s)-k]
	}
    return nil
}

// getKthFromEndV2 双指针法
// 前指针和后指针：前指针先走K步，后指针再和前指针一起走，当前指针到尾节点跳出后后指针就是第K个节点（两个指针的距离K-1）
func getKthFromEndV2(head *ListNode, k int) *ListNode {
	former, latter := head, head

	for i := 0 ; i < k; i++ {
		former = former.Next
	}

    for former != nil {
		former, latter := former.Next, latter.Next
	}
    return latter
}