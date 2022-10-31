/**
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

 此外，你可以假设该网格的四条边均被水包围。



 示例 1：


输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1


 示例 2：


输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3




 提示：


 m == grid.length
 n == grid[i].length
 1 <= m, n <= 300
 grid[i][j] 的值为 '0' 或 '1'


 Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 1966 👎 0

*/

package leetcode

func Q200Main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	islandNum := 0
	for row, rowLine := range grid {
		for col, isLand := range rowLine {
			if isLand == '1' {
				islandNum++
				grid = dfs(grid, row, col)
			}
		}
	}
	return islandNum
}

func dfs(image [][]byte, sr int, sc int) [][]byte {
	image[sr][sc] = '0'
	// 右
	if sc+1 < len(image[sr]) && image[sr][sc+1] == '1' {
		image = dfs(image, sr, sc+1)
	}
	// 下
	if sr+1 < len(image) && image[sr+1][sc] == '1' {
		image = dfs(image, sr+1, sc)
	}
	// 左
	if sc-1 >= 0 && image[sr][sc-1] == '1' {
		image = dfs(image, sr, sc-1)
	}
	// 上
	if sr-1 >= 0 && image[sr-1][sc] == '1' {
		image = dfs(image, sr-1, sc)
	}
	return image
}

// leetcode submit region end(Prohibit modification and deletion)
