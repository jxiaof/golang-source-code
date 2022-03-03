from typing import List

# Definition for singly-linked list.


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def getKthFromEnd(self, head: ListNode, k: int) -> ListNode:
        if not head:
            return None
        if not head.next:
            return head
        p1 = head
        p2 = head
        for _ in range(k):
            p2 = p2.next
        while p2:
            p1 = p1.next
            p2 = p2.next
        return p1


def foo(head: ListNode, k: int) -> ListNode:
    if not head:
        return None
    if not head.next:
        return head
    a, b = head, head
    for _ in range(k):
        b = b.next
    while b:
        a, b = a.next, b.next
    return a

if __name__ == "__main__":
    li = [1, 2, 3, 4, 5]
    j = 2
    head = ListNode(li[0])
    p = head
    for i in li[1:]:
        p.next = ListNode(i)
        p = p.next
    s = Solution()
    n = s.getKthFromEnd(head, j)
    for _ in range(j):
        print(n.val)
        n = n.next
    print("\n")
    n = foo(head, j)
    for _ in range(j):
        print(n.val)
        n = n.next

# Output:



