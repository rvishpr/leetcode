/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
 
int inorder(struct TreeNode *node, int *count, int k) {
    if (node == NULL) return -1;
    int val;
    
    val = inorder(node->left, count, k);

    (*count)++;

    if (*count == k) {
        return node->val;
    } else if (*count > k) return val;
   
    val = inorder(node->right, count, k);

    if (*count >= k) return val;
    
    return -1;
}

int kthSmallest(struct TreeNode* root, int k) {
    int count = 0;
    return inorder(root, &count, k);
}
