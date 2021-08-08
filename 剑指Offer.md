#### 100. 三角形中最小路径之和

**题目链接：**[100. 三角形中最小路径之和](https://leetcode-cn.com/problems/IlPe0q/)

**题目描述：**

给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。**相邻的结点**在这里指的是**下标** 与 **上一层结点下标** 相同或者等于 **上一层结点下标 + 1**的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。

**解题思路：**从题目的描述来看，第i层的数据是由第i-1层推导而来，自然而然想到这道题目非常适合使用动态规划来求解。

**解题步骤：**

本题使用动态规划求解，下面我按照动态规划的求解步骤进行分析。

1、确定动态规划数组dp\[i]\[j]的含义

在这道题目中dp\[i]\[j]表示到达第i层第j个位置的最小路径和。

2，确定状态转移方程

能够影响dp\[i]\[j]的只有dp\[i-1]\[j-1]和dp\[i-1]\[j]（根据题目中描述的相邻节点得出）

3、dp数组初始化

dp数组只需要初始化第0行，dp[0] = {triangle[0]}

4、遍历顺序

题目描述的移动方向是从上而下，那么遍历顺序自然也是从上而下（dp\[i]\[j]由dp\[i-1]\[j]和dp\[i-1]\[j-1]推导而来）。

5、分析dp数组的计算过程

以triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]为输入进行分析

dp初始化状态：

索引位置  	0		1		2		3		4

​          i=0      2

​          i=1      5    	6

​          i=2      11 	10      13

​          i=3      15      **11**      18       16

6、输出结果

题目要求输出到达最低层的最小值，因此需要输出最底层数组的最小值。



**源码分析：**

~~~go
func MinimumTotal(triangle [][]int) int {
    // 特殊输入处理
	if len(triangle) == 0 {
		return 0
	}

    // 创建dp二维数组
	rows := len(triangle)
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
    
    // 初始化dp数组
	dp[0][0] = triangle[0][0]
    
    // 遍历
	for i := 1; i < rows; i++ {
        // 针对每行的第一个和最后一个位置的计算特殊处理
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		dp[i][len(triangle[i])-1] = dp[i-1][len(triangle[i])-2] + triangle[i][len(triangle[i])-1]
		
        // 按照状态转移方程计算
        for j := 1; j < len(triangle[i])-1; j++ {
			dp[i][j] = leetcode_go.Min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
	}

    // 输出最小值
	return leetcode_go.MinValueOfSlice(dp[rows-1])

}
~~~

