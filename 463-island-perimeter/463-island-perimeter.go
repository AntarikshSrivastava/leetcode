func islandPerimeter(grid [][]int) int {
    
    var res int
    
    for i:=0; i<len(grid); i++ {
        
        for j:=0; j<len(grid[0]); j++ {
            
            if grid[i][j] == 1 {
                res += 4
                
                if i>0 && grid[i-1][j] == 1 {
                    res -= 2
                }
                
                if j>0 && grid[i][j-1] == 1 {
                    res -= 2
                }
            }
        }
    }
    return res
}