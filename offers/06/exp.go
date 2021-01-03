// 从头到尾打印链表
// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // reversePrint 递归实现
 func reversePrint(head *ListNode) []int {
	 
    nodes := make([]int,0)
    if head == nil {
        return nodes
    }
    return append(reversePrint(head.Next), head.Val)
}

// reversePrintV2 先反转链表后打印
func reversePrintV2(head *ListNode)[]int {
	if head == nil {
        return nil
	}
	// 反转链表
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	nodes := make([]int,0)
	for pre != nil {
        nodes = append(nodes, pre.Val)
        pre = pre.Next
	}
	return nodes
}

// reversePrintV3 打印存储到数组中，反转数组
func reversePrintV3(head *ListNode) []int {

	nodes := make([]int,0)
    for head != nil {
        nodes = append(nodes, head.Val)
        head = head.Next
    }
   
    for i, j := 0, len(nodes)-1; i < j; {
        nodes[i], nodes[j] = nodes[j], nodes[i]
        i++
        j--
    }
    return nodes
}