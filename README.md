# 数据结构及算法的Go语言实现

在学习数据结构和算法的过程中,根据自己的理解,用Go实现一遍代码.
推荐一本学习数据结构的书:

![](img/data.png)

讲的浅显易懂,语言也很幽默风趣

再推荐一个数据结构和算法的网站 [Data Structure Visualizations](https://www.cs.usfca.edu/~galles/visualization/Algorithms.html)

这个网站里面包含了很多算法,以及动画实现的过程,例如排序:

![](img/sort.png)

例如队列的添加删除过程

![Alt text](img/queue.png)

> 刷题必备网站推荐:
[leetcode](https://leetcode.com/explore/)
[leetcode中文版](https://leetcode-cn.com/explore/)
北大ACM, 网上找相关资料

## 一. 数据结构源码

### 1. 顺序结构线性链表
自己实现了一遍顺序结构链表,写在list.go中,list_test.go是测试代码


## 二. 数据结构相关面试题
面试的时候有一个题目:
> 1.给定两个非空链表来表示两个非负整数,位数按照逆序方式存储,他们的每个节点只存储单个数字,将两数相加返回一个新的链表.
你可以假设除了数字0之外,这两个数字不会以0开头.
**示例:**
输入:(2->4->3) + (5->6->4)
输出:7->0->8
原因:342+465 = 807

我自己实现了一遍,写在example/addlist.go 里面

> 2.合并 k 个排序链表，返回合并后的排序链表。
```
输入:
[
    1->4->5,
    1->3->4,
    2->6
]

 输出: 1->1->2->3->4->4->5->6
 ```

