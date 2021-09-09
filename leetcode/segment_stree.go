package main

import (
	"fmt"
)

type SegmentTreeNode struct {
	data   int
	start  int
	end    int
	marked bool
	iv     int
}

type SegmentTree struct {
	nodes []SegmentTreeNode
	data  []int
}

func (tree SegmentTree) Init() {
	tree.build(0, len(tree.data)-1, 1)
}

func (tree SegmentTree) build(begin, end int, p int) {
	// 如果区间长度为1，则建数结束
	if begin == end {
		node := SegmentTreeNode{
			data:   tree.data[begin],
			start:  begin,
			end:    end,
			marked: false,
			iv:     0,
		}
		tree.nodes[p] = node
		return
	}

	// 划分区间
	// 移位运算符的优先级小于加减法，所以加上括号
	// 如果写成 (begin + end) >> 1 可能会超出 int 范围
	m := begin + ((end - begin) >> 1)

	// 递归对左右区间建树
	tree.build(begin, m, p*2)
	tree.build(m+1, end, p*2+1)

	node := SegmentTreeNode{
		data:   tree.nodes[p*2].data + tree.nodes[p*2+1].data,
		start:  begin,
		end:    end,
		marked: false,
		iv:     0,
	}
	tree.nodes[p] = node
}

func (tree SegmentTree) GetRangeSumWithNoModify(left, right int) int {
	return tree.getRangeSum(left, right, 1, len(tree.data), 1)
}

// 数据没有修改的条件下区间查询
func (tree SegmentTree) getRangeSum(l, r, s, t, p int) int {
	// [l,r] 为查询区间,[s,t] 为当前节点包含的区间,p 为当前节点的编号
	if l <= s && t <= r {
		return tree.nodes[p].data // 当前区间为询问区间的子集时直接返回当前区间的和
	}
	m := s + ((t - s) >> 1)
	sum := 0
	if l <= m {
		sum += tree.getRangeSum(l, r, s, m, p*2)
	}
	// 如果左儿子代表的区间 [l,m] 与询问区间有交集,则递归查询左儿子
	if r > m {
		sum += tree.getRangeSum(l, r, m+1, t, p*2+1)
	}
	// 如果右儿子代表的区间 [m+1,r] 与询问区间有交集,则递归查询右儿子
	return sum
}

func (tree SegmentTree) ModifyRange(l, r, c int) {
	tree.modifyRange(l, r, c, 1, len(tree.data), 1)
}

func (tree SegmentTree) modifyRange(l, r, c, s, t, p int) {
	// [l,r] 为修改区间,c 为被修改的元素的变化量,[s,t] 为当前节点包含的区间,p
	// 为当前节点的编号
	if l <= s && t <= r {
		tree.nodes[p].data = tree.nodes[p].data + (t-s+1)*c
		tree.nodes[p].marked = true
		tree.nodes[p].iv = tree.nodes[p].iv + c
		return
	} // 当前区间为修改区间的子集时直接修改当前节点的值,然后打标记,结束修改

	m := s + ((t - s) >> 1)
	if tree.nodes[p].iv != 0 && s != t {
		// 如果当前节点的懒标记非空,则更新当前节点两个子节点的值和懒标记值
		tree.nodes[2*p].data += tree.nodes[p].iv * (m - s + 1)

		tree.nodes[2*p+1].data += tree.nodes[p].iv * (t - m)

		tree.nodes[2*p].iv += tree.nodes[p].iv

		tree.nodes[p*2+1].iv += tree.nodes[p].iv // 将标记下传给子节点

		// 清空根节点
		tree.nodes[p].marked = false
		tree.nodes[p].iv = 0
	}
	if l <= m {
		tree.modifyRange(l, r, c, s, m, p*2)
	}
	if r > m {
		tree.modifyRange(l, r, c, m+1, t, p*2+1)
	}

	tree.nodes[p].data = tree.nodes[p*2].data + tree.nodes[p*2+1].data
}

func (tree SegmentTree) GetSumWhenModify(l, r int) int {
	return tree.getSumWhenModify(l, r, 0, len(tree.data), 1)
}

func (tree SegmentTree) getSumWhenModify(l, r, s, t, p int) int {
	// [l,r] 为查询区间,[s,t] 为当前节点包含的区间,p为当前节点的编号
	if l <= s && t <= r {
		return tree.nodes[p].data
	}
	// 当前区间为询问区间的子集时直接返回当前区间的和
	m := s + ((t - s) >> 1)
	if tree.nodes[p].iv != 0 {
		// 如果当前节点的懒标记非空,则更新当前节点两个子节点的值和懒标记值
		tree.nodes[p*2].data += tree.nodes[p].iv * (m - s + 1)
		tree.nodes[p*2+1].data += tree.nodes[p].iv * (t - m)

		tree.nodes[p*2].iv = tree.nodes[p].iv
		tree.nodes[p*2+1].iv = tree.nodes[p].iv // 将标记下传给子节点
		tree.nodes[p].iv = 0                    // 清空当前节点的标记
		tree.nodes[p].marked = false
	}
	sum := 0
	if l <= m {
		sum = tree.getSumWhenModify(l, r, s, m, p*2)
	}
	if r > m {
		sum += tree.getSumWhenModify(l, r, m+1, t, p*2+1)
	}
	return sum
}

func (tree SegmentTree) show() {
	for i := 0; i < len(tree.nodes); i++ {
		fmt.Printf("%d ", tree.nodes[i].data)
	}
	fmt.Print("\n")
}

func main() {
	//a := []int{10, 11, 12, 13, 14}
	//d := make([]int, 20)
	//b := make([]int, 20)
	//buildSegmentTree(0, 4, 1, &d, a)
	//
	//modifyRange(3, 5, 1, 1, 5, 1, d, b)
	//result := getSum(3, 5, 1, 5, 1, d, b)
	//fmt.Println(result)
	tree := SegmentTree{
		data:  []int{10, 11, 12, 13, 14},
		nodes: make([]SegmentTreeNode, 20),
	}

	tree.Init()

	fmt.Println(tree.GetRangeSumWithNoModify(3, 5))

}

// 对区间[begin,end]建立线段树，当前根节点的编号为p，st表示线段树，a表示原始输入
func buildSegmentTree(begin, end int, p int, st *[]int, a []int) {
	// 如果区间长度为1，则建数结束
	if begin == end {
		(*st)[p] = a[begin]
		return
	}

	// 划分区间
	// 移位运算符的优先级小于加减法，所以加上括号
	// 如果写成 (begin + end) >> 1 可能会超出 int 范围
	m := begin + ((end - begin) >> 1)

	// 递归对左右区间建树
	buildSegmentTree(begin, m, p*2, st, a)
	buildSegmentTree(m+1, end, p*2+1, st, a)

	(*st)[p] = (*st)[p*2] + (*st)[(p*2)+1]
}

func getRangeSum(l, r, s, t, p int, d []int) int {
	// [l,r] 为查询区间,[s,t] 为当前节点包含的区间,p 为当前节点的编号
	if l <= s && t <= r {
		return d[p] // 当前区间为询问区间的子集时直接返回当前区间的和
	}
	m := s + ((t - s) >> 1)
	sum := 0
	if l <= m {
		sum += getRangeSum(l, r, s, m, p*2, d)
	}
	// 如果左儿子代表的区间 [l,m] 与询问区间有交集,则递归查询左儿子
	if r > m {
		sum += getRangeSum(l, r, m+1, t, p*2+1, d)
	}
	// 如果右儿子代表的区间 [m+1,r] 与询问区间有交集,则递归查询右儿子
	return sum
}

func modifyRange(l, r, c, s, t, p int, d []int, b []int) {
	// [l,r] 为修改区间,c 为被修改的元素的变化量,[s,t] 为当前节点包含的区间,p
	// 为当前节点的编号
	if l <= s && t <= r {
		d[p] += (t - s + 1) * c
		b[p] += c
		return
	} // 当前区间为修改区间的子集时直接修改当前节点的值,然后打标记,结束修改

	m := s + ((t - s) >> 1)
	if b[p] != 0 && s != t {
		// 如果当前节点的懒标记非空,则更新当前节点两个子节点的值和懒标记值
		d[p*2] += b[p] * (m - s + 1)
		d[p*2+1] += b[p] * (t - m)
		b[p*2] += b[p]
		b[p*2+1] += b[p] // 将标记下传给子节点
		b[p] = 0         // 清空当前节点的标记
	}
	if l <= m {
		modifyRange(l, r, c, s, m, p*2, d, b)
	}
	if r > m {
		modifyRange(l, r, c, m+1, t, p*2+1, d, b)
	}
	d[p] = d[p*2] + d[p*2+1]
}

func getSum(l, r, s, t, p int, d []int, b []int) int {
	// [l,r] 为查询区间,[s,t] 为当前节点包含的区间,p为当前节点的编号
	if l <= s && t <= r {
		return d[p]
	}
	// 当前区间为询问区间的子集时直接返回当前区间的和
	m := s + ((t - s) >> 1)
	if b[p] != 0 {
		// 如果当前节点的懒标记非空,则更新当前节点两个子节点的值和懒标记值
		d[p*2] += b[p] * (m - s + 1)
		d[p*2+1] += b[p] * (t - m)
		b[p*2] += b[p]
		b[p*2+1] += b[p] // 将标记下传给子节点
		b[p] = 0         // 清空当前节点的标记
	}
	sum := 0
	if l <= m {
		sum = getSum(l, r, s, m, p*2, d, b)
	}
	if r > m {
		sum += getSum(l, r, m+1, t, p*2+1, d, b)
	}
	return sum
}
