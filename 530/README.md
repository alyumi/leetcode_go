# 530. Minimum Absolute Difference in BST


Difficulty: <span style="color:green">EASY</span>

## Problem
Given the root of a Binary Search Tree (BST), return the minimum absolute difference between the values of any two different nodes in the tree.

 

### Example 1:

```mermaid
graph TD;
    4-->2;
    4-->6;
    2-->1;
    2-->3;
```

Input: root = [4,2,6,1,3]
Output: 1


### Example 2:
```mermaid
graph TD;
    1-->0;
    1-->48;
    48-->12;
    48-->49;
```

Input: root = [1,0,48,null,null,12,49]
Output: 1
 

### Constraints:

The number of nodes in the tree is in the range [2, 10^4].
0 <= Node.val <= 10^5
 

Note: This question is the same as 783: https://leetcode.com/problems/minimum-distance-between-bst-nodes/