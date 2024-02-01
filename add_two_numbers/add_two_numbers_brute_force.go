import "fmt"

// type ListNode struct {
//     Val int
//     Next *ListNode
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var firstNode *ListNode
	var newNode *ListNode
	var aux *ListNode
	l1Walker, l2Walker, currentL1Value, currentL2Value, quotient, sum := l1, l2, 0, 0, 0, 0

	for l1Walker != nil || l2Walker != nil {

		if l1Walker != nil {
			currentL1Value = l1Walker.Val
			l1Walker = l1Walker.Next
		} else {
			currentL1Value = 0
		}

		if l2Walker != nil {
			currentL2Value = l2Walker.Val
			l2Walker = l2Walker.Next
		} else {
			currentL2Value = 0
		}

		if quotient > 0 {
			sum = currentL1Value + currentL2Value + quotient
		} else {
			sum = currentL1Value + currentL2Value
		}

		if firstNode == nil {
			firstNode = &ListNode{sum % 10, nil}
			newNode = firstNode
			aux = firstNode
		} else {
			newNode = &ListNode{sum % 10, nil}
			aux.Next = newNode
			aux = aux.Next
		}

		quotient = sum / 10

	}

	if quotient > 0 {
		newNode = &ListNode{quotient, nil}
		aux.Next = newNode
	}

	return firstNode

}