package easy

import (
	"testing"
)

func TestMergeTwoLists2(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2, Next: &ListNode{Val: 3}}

	mergeTwoListsWrong(l1, l2)
	mergeTwoLists(l1, l2)
	mergeTwoListsOptimization(l1, l2)
}
