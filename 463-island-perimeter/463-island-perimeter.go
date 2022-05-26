func islandPerimeter(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    ans := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            tmp := 0
            if grid[i][j] == 1 {
                if i == 0 || grid[i - 1][j] == 0 { tmp++ }
                if i == m - 1 || grid[i + 1][j] == 0 { tmp++ }
                if j == 0 || grid[i][j - 1] == 0 { tmp++ }
                if j == n - 1 || grid[i][j + 1] == 0 { tmp++ }
            }
            ans += tmp
        }
    }
    return ans
}