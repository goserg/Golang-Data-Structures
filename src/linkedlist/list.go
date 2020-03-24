package linkedlist

import (
	"fmt"
	"strconv"
)

//LinkedList is a linear collection of data elements
type LinkedList struct {
	head node
	size int
}

//NewLinkedList returns empty linked list
func NewLinkedList() *LinkedList {
	ll := new(LinkedList)
	ll.head = node{
		next:  nil,
		value: 0}
	ll.size = 0
	return ll
}

//AddFirst inserts data to the begining of the list
func (ll *LinkedList) AddFirst(data int) {
	newNode := node{
		next:  ll.head.next,
		value: data,
	}
	ll.head.next = &newNode
	ll.size++
}

//Append inserts data to the end of the list
func (ll *LinkedList) Append(data int) {
	if ll.size == 0 {
		ll.AddFirst(data)
		return
	}
	newNode := node{
		next:  nil,
		value: data,
	}
	currentNode := ll.head.next
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = &newNode
	ll.size++
}

//Get returns data
func (ll LinkedList) Get(index int) (int, error) {
	if ll.size == 0 {
		return 0, fmt.Errorf("Pop! List is empty")
	}
	if index < 0 || index > ll.size-1 {
		return 0, fmt.Errorf(fmt.Sprintf("Get Index %d is out of bounds", index))
	}
	currentNode := ll.head.next
	count := 0
	for count < index {
		currentNode = currentNode.next
		count++
	}
	return currentNode.value, nil
}

//Pop deletes data from lisr and returns it
func (ll *LinkedList) Pop(index int) (int, error) {
	if ll.size == 0 {
		return 0, fmt.Errorf("Pop! List is empty")
	}
	if index < 0 || index > ll.size-1 {
		return 0, fmt.Errorf(fmt.Sprintf("Pop Index %d is out of bounds", index))
	}
	currentNode := ll.head.next
	lastNode := &ll.head
	count := 0
	for count < index {
		lastNode = currentNode
		currentNode = currentNode.next
		count++
	}
	lastNode.next = currentNode.next
	ll.size--
	return currentNode.value, nil
}

//Search returns index of the data, if data is not on the list return -1
func (ll LinkedList) Search(data int) int {
	if ll.size == 0 {
		return -1
	}
	currentNode := ll.head.next
	index := 0
	for currentNode != nil {
		if currentNode.value == data {
			return index
		}
		currentNode = currentNode.next
		index++
	}
	return -1
}

//Clear erasing all data from list
func (ll *LinkedList) Clear() {
	ll.head = node{
		next:  nil,
		value: 0}
	ll.size = 0
}

func (ll LinkedList) String() string {
	if ll.size == 0 {
		return "[]"
	}
	currentNode := ll.head.next
	data := strconv.Itoa(currentNode.value)
	for currentNode.next != nil {
		currentNode = currentNode.next
		data = data + ", " + strconv.Itoa(currentNode.value)
	}
	return "[" + data + "]"
}

//GetSize returns size of the linked list
func (ll LinkedList) GetSize() int {
	return ll.size
}
