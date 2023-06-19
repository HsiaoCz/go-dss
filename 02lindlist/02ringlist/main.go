package main

// 循环链表
type Ring struct {
	next  *Ring       // 后继节点
	prev  *Ring       // 前驱节点
	Value interface{} // 数据
}

// 初始化空的循环链表
// 空的循环链表前驱和后继都指向自己
// 这里的时间复杂度为O(1)
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// 创建N个节点的循环链表
// 这里时间复杂度为O(n)
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// 获取下一个节点
// 获取前驱和后继时间复杂度为O(1)
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

// 获取第n个节点
// 循环链表有前驱和后继，所以如果n为负数
// 表示往前面遍历，否则表示往后面遍历
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}
func main() {
	r := new(Ring)
	r.init()
}
