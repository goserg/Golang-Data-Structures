package circularlist

import (
	"fmt"
)

//CLL is a circular collection of data elements
type CLL struct {
	tail node
	size int
}

//NewCLL returns empty circular linked list
func NewCLL() *CLL {
	cll := new(CLL)
	cll.tail = node{
		next:  nil,
		value: 0}
	cll.size = 0
	return cll
}

//AddFirst inserts data to the begining of the list
func (cll *CLL) AddFirst(data interface{}) {
	newNode := node{
		next:  nil,
		value: data,
	}
	if cll.size == 0 {
		newNode.next = &newNode
		cll.tail.next = &newNode
	} else {
		newNode.next = cll.tail.next.next
		cll.tail.next.next = &newNode
	}
	cll.size++
}

//Append inserts data to the end of the list
func (cll *CLL) Append(data interface{}) {
	if cll.size == 0 {
		cll.AddFirst(data)
		return
	}
	temp := cll.tail.next
	newNode := node{
		next:  cll.tail.next.next,
		value: data,
	}
	cll.tail.next = &newNode
	temp.next = &newNode
	cll.size++
}

//Get returns data
func (cll CLL) Get(index int) (interface{}, error) {
	if cll.size == 0 {
		return 0, fmt.Errorf("Pop! List is empty")
	}
	if index < 0 || index > cll.size-1 {
		return 0, fmt.Errorf(fmt.Sprintf("Get Index %d is out of bounds", index))
	}
	currentNode := cll.tail.next.next
	count := 0
	for count < index {
		currentNode = currentNode.next
		count++
	}
	return currentNode.value, nil
}

// Set changes data on specified position.
// Returns error
func (cll *CLL) Set(index int, data int) error {
	if cll.size == 0 {
		return fmt.Errorf("Set fail! List is empty")
	}
	if index < 0 || index > cll.size-1 {
		return fmt.Errorf(fmt.Sprintf("Set fail! Index %d is out of bounds", index))
	}
	currentNode := cll.tail.next.next
	count := 0
	for count < index {
		currentNode = currentNode.next
		count++
	}
	currentNode.value = data
	return nil
}

//Pop deletes data from lisr and returns it
func (cll *CLL) Pop(index int) (interface{}, error) {
	if cll.size == 0 {
		return 0, fmt.Errorf("Pop! List is empty")
	}
	if index < 0 || index > cll.size-1 {
		return 0, fmt.Errorf(fmt.Sprintf("Pop Index %d is out of bounds", index))
	}
	currentNode := cll.tail.next.next
	lastNode := cll.tail.next
	count := 0
	for count < index {
		lastNode = currentNode
		currentNode = currentNode.next
		count++
	}
	lastNode.next = currentNode.next
	cll.size--
	return currentNode.value, nil
}

//Search returns index of the data, if data is not on the list return -1
func (cll CLL) Search(data interface{}) int {
	if cll.size == 0 {
		return -1
	}
	currentNode := cll.tail.next.next
	if currentNode.value == data {
		return 0
	}
	index := 0
	for currentNode != cll.tail.next {
		index++
		currentNode = currentNode.next
		if currentNode.value == data {
			return index
		}
	}
	return -1
}

//Clear erasing all data from list
func (cll *CLL) Clear() {
	cll.tail = node{
		next:  nil,
		value: 0}
	cll.size = 0
}

func (cll CLL) String() string {
	if cll.size == 0 {
		return "[]"
	}
	var currentNode *node
	currentNode = cll.tail.next.next
	data := fmt.Sprintf("%v", currentNode.value)
	for currentNode.next != cll.tail.next.next {
		currentNode = currentNode.next
		data = data + ", " + fmt.Sprintf("%v", currentNode.value)
	}
	return "[" + data + "]"
}

//GetSize returns size of the linked list
func (cll CLL) GetSize() int {
	return cll.size
}
