package lru

type node struct {
	next  *node
	prev  *node
	key   string
	value string
}

type doublyLinkedList struct {
	head   *node
	tail   *node
	length int
}

func NewList() *doublyLinkedList {
	return &doublyLinkedList{}
}

// the node is always added at the head
func (list *doublyLinkedList) addNode(node *node) {
	list.length++

	// in case there is no node in the doubly linked list
	if list.head == nil {
		list.head = node
		list.tail = node

		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.next = list.head
	list.head.prev = node
	node.prev = nil
	list.head = node
	// free the node pointer
	node = nil

	return
}

// remove a specific node or the tail
func (list *doublyLinkedList) removeNode(node *node) {
	list.length--

	// update the tail pointer to point to the previous node
	if list.tail == node {
		list.tail = node.prev
	}
	if list.head == node {
		list.head = node.next
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	// free the node pointer
	node = nil

	return
}
