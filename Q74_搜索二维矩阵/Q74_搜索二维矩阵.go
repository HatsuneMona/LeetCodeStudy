/*
74. 搜索二维矩阵
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。
该矩阵具有如下特性：
	每行中的整数从左到右按升序排列。
	每行的第一个整数大于前一行的最后一个整数。

示例 1：
	1  3  5  7
	10 11 16 20		输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
	23 30 34 60		输出：true

示例 2：
	1  3  5  7
	10 11 16 20		输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
	23 30 34 60		输出：false

提示：
	m == matrix.length
	n == matrix[i].length
	1 <= m, n <= 100
	-104 <= matrix[i][j], target <= 104
*/
package main

func main() {

}

/*
My Note：
	可以用两次二分查找法，第一次大概查找到它是哪一行，第二次从具体的一行中找到
*/

func searchMatrix(matrix [][]int, target int) bool {
	if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	whichCol, where := 0, 0
	maxInCol := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		maxInCol[i] = matrix[i][len(matrix[0])-1]
	}
	whichCol = BinarySearch(maxInCol, target, 0, len(maxInCol)-1, false)
	if whichCol >= 0 {
		where = BinarySearch(matrix[whichCol], target, 0, len(matrix[whichCol])-1, true)
	}
	if where >= 0 {
		return true
	} else {
		return false
	}
}

func BinarySearch(matrix []int, target, a, b int, precise bool) int {
	if b == a+1 || b == a-1 || a == b { //二分查找该结束了
		if precise == true { //需要精确查找（在行中找值的时候）
			if target == matrix[a] { //找到目标（精确）
				return a
			} else if target == matrix[b] {
				return b
			}
			return -1
		} else { //不需要精确查找（在找列的时候）
			if target <= matrix[a] {
				return a
			} else {
				return b
			}
		}
	} else if target < matrix[(b+a)/2] { //目标在前半部分
		return BinarySearch(matrix, target, a, (b+a)/2, precise)
	} else if target > matrix[(b+a)/2] { //目标在后半部分
		return BinarySearch(matrix, target, (b+a)/2, b, precise)
	} else if target == matrix[(b+a)/2] { //找到目标（精确）
		return (b + a) / 2
	}
	return -2
}
