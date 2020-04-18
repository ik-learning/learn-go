package main

// import (
// 	"bufio"
// 	"fmt"
// 	"testing"

// 	assert "github.com/stretchr/testify/assert"
// )

// type SinglyLinkedListNode struct {
// 	data int32
// 	next *SinglyLinkedListNode
// }

// type SinglyLinkedList struct {
// 	head *SinglyLinkedListNode
// 	tail *SinglyLinkedListNode
// }

// func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
// 	node := &SinglyLinkedListNode{
// 		next: nil,
// 		data: nodeData,
// 	}

// 	if singlyLinkedList.head == nil {
// 		singlyLinkedList.head = node
// 	} else {
// 		singlyLinkedList.tail.next = node
// 	}

// 	singlyLinkedList.tail = node
// }

// func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
// 	for node != nil {
// 		fmt.Fprintf(writer, "%d", node.data)

// 		node = node.next

// 		if node != nil {
// 			fmt.Fprintf(writer, sep)
// 		}
// 	}
// }

// type SinglyLinkedListNode struct {
// 	data int32
// 	next *SinglyLinkedListNode
// }

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
// func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {

// }

// func TestNodeAtPosition(t *testing.T) {
	// var tests = []struct {
	// 	first    string
	// 	second   string
	// 	expected string
	// }{
	// 	{"ab", "abc", "YES"},
	// 	{"hi", "world", "NO"},
	// }

	// for _, test := range tests {
	// 	result := insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32)(test.first, test.second)
	// 	assert.Equal(t, test.expected, result, "they should be equal")
	// }
// }
