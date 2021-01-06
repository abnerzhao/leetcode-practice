// LRU最近最少使用内存淘汰算法
/*
思路：
1.单链表实现
越靠近链表尾部是越早之前访问的数据。当一个新数据被访问时，从头顺序遍历链表。
- 如果数据在链表中，则将其从原来位置删除，再插入到链表头部
- 如果数据不存在链表中
  - 缓存未满，则将节点直接插入到链表头部
  - 缓存已满，则将链表尾部删除，将新数据插入链表头部
*/
package main

// LinkNode 单链表
type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

// Cache 缓存模型
type Cache struct {
	Head *LinkNode // 头节点
	Max int64  // 缓存最大长度
}


// Insert 缓存写入数据
func(c *Cache) Insert(newNode interface{}) {
	// 头部为空则直接插入头节点
	if c.Head == nil {
		c.Head = &LinkNode{
			Next: nil,
			Data: newNode,
		}
		return
	}

	var pre *LinkNode
	current := c.Head
	i := 1

	n := &LinkNode{
		Next: c.Head,
		Data: NewNode,
	}


	for {
		if current.Data == newNode {
			if pre != nil {
				pre.Next = current.Next
				c.Head = n
			}
			break
		}

		if current.Next == nil {
			if i >= c.Max {
				pre.Next = nil
				c.Head = n
			} else {
				c.Head = n
			}
			break
		}

		pre = current
		current = current.Next
		i++
	}
}
