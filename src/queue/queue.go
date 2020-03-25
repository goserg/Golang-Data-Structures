package queue

import "fmt"

type node struct {
	next  *node
	value interface{}
}

// Queue is a collection of entities that are maintained in a sequence
// and can be modified by the addition of entities at one end of the sequence
// and removal from the other end of the sequence.
type Queue struct {
	tail node
	size int
}

// NewQueue is a constructor for Queue
func NewQueue() *Queue {
	q := new(Queue)
	q.tail = node{
		next:  nil,
		value: 0}
	q.size = 0
	return q
}

//Push inserts data to the Queue
func (q *Queue) Push(data interface{}) {
	newNode := node{
		next:  nil,
		value: data,
	}
	if q.size == 0 {
		newNode.next = &newNode
		q.tail.next = &newNode
		q.size++
		return
	}
	temp := q.tail.next
	newNode.next = q.tail.next.next
	q.tail.next = &newNode
	temp.next = &newNode
	q.size++
}

//Peek returns data from top of the Queue
func (q Queue) Peek() (interface{}, error) {
	if q.size == 0 {
		return 0, fmt.Errorf("Pop! List is empty")
	}
	return q.tail.next.next.value, nil
}

//Pop deletes data from Queue and returns it
func (q *Queue) Pop() (interface{}, error) {
	if q.size == 0 {
		return 0, fmt.Errorf("Pop! List is empty")
	}
	value := q.tail.next.next.value
	q.tail.next.next = q.tail.next.next.next
	q.size--
	return value, nil
}

//Search returns index of the data, if data is not on the Queue return -1.
// Index starts from 1
func (q Queue) Search(data interface{}) int {
	if q.size == 0 {
		return -1
	}
	currentNode := q.tail.next.next
	if currentNode.value == data {
		return 1
	}
	index := 1
	for currentNode != q.tail.next {
		index++
		currentNode = currentNode.next
		if currentNode.value == data {
			return index
		}
	}
	return -1
}

//Clear erasing all data from Queue
func (q *Queue) Clear() {
	q.tail = node{
		next:  nil,
		value: 0}
	q.size = 0
}

func (q Queue) String() string {
	if q.size == 0 {
		return "[]"
	}
	var currentNode *node
	currentNode = q.tail.next.next
	data := fmt.Sprintf("%v", currentNode.value)
	for currentNode.next != q.tail.next.next {
		currentNode = currentNode.next
		data = data + ", " + fmt.Sprintf("%v", currentNode.value)
	}
	return "[" + data + "]"
}

//GetSize returns size of the Queue
func (q Queue) GetSize() int {
	return q.size
}

//IsEmpty returns true if Queue is empty, else false
func (q Queue) IsEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}
