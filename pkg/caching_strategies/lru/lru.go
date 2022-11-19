package lru

import (
	"errors"
	"fmt"
)

type LRU struct {
	size  int
	cache map[string]*node
	list  *doublyLinkedList
}

func NewLRU(size int) (*LRU, error) {
	if size <= 0 {
		return nil, errors.New("size cannot be less than or equal to 0")
	}
	return &LRU{
		size:  size,
		cache: make(map[string]*node, 0),
		list:  NewList(),
	}, nil
}

func (lru *LRU) Get(key string) (value string, ok bool) {
	if _, ok := lru.cache[key]; !ok {
		return "", false
	}
	nodeTemp := lru.cache[key]
	lru.list.removeNode(lru.cache[key])
	lru.list.addNode(nodeTemp)
	lru.cache[key] = nodeTemp

	return lru.cache[key].value, true
}

func (lru *LRU) Delete(key string) bool {

	if _, ok := lru.cache[key]; !ok {
		return false
	}
	lru.list.removeNode(lru.cache[key])
	delete(lru.cache, key)

	return true
}

func (lru *LRU) Set(key string, value string) bool {

	if lru.list.length == lru.size {
		nodeTemp := lru.list.tail
		lru.list.removeNode(nodeTemp)
		delete(lru.cache, nodeTemp.key)
	}

	nodeTemp := &node{
		key:   key,
		value: value,
	}
	if _, ok := lru.cache[key]; ok {
		nodeTemp = lru.cache[key]
		lru.list.removeNode(lru.cache[key])
	}

	// if the element is not in the map
	lru.list.addNode(nodeTemp)
	lru.cache[key] = nodeTemp

	return true
}

func (lru *LRU) Len() int {

	return lru.size
}

func (lru *LRU) Print() {
	fmt.Println("---moving forwards----")
	head := lru.list.head
	for head != nil {
		fmt.Println(head.key + "----" + head.value)
		head = head.next
	}

	fmt.Println("---moving backwards----")

	tail := lru.list.tail
	for tail != nil {
		fmt.Println(tail.key + "----" + tail.value)
		tail = tail.prev
	}
}
