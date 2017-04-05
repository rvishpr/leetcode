/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
 
typedef struct ListNode ListNode;
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
    int sum, carry = 0;
    ListNode *node = NULL;
    ListNode *root = NULL;
    
    while (l1 != NULL || l2 != NULL) {
        ListNode *node1 = (ListNode *)malloc(sizeof(ListNode));
        if (node1 == NULL) {
            return NULL;//free everything TODO
        }
        
        sum = carry;
        if (l1) {
            sum += l1->val;
            l1 = l1->next;
        }        
        
        if (l2) {
            sum += l2->val;
            l2 = l2->next;
        }
        
        node1->val = sum%10;
        carry = sum/10;
        node1->next = NULL;
        
        if (root == NULL) {
            root = node1;
            node = node1;
        } else {
            node->next = node1;
            node = node1;
        }
    }
    if (carry != 0) {
        ListNode *node1 = (ListNode *)malloc(sizeof(ListNode));
        if (node1 == NULL) {
            return NULL;//free everything TODO
        }
        
        node1->val = carry;
        node1->next = NULL;
        
        if (root == NULL) {
            root = node1;
        } else {
            node->next = node1;
        }
    }
    return root;
}
