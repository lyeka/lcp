package medium

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
	}}

	after := removeNthFromEnd(head, 2)
	print(after)
}