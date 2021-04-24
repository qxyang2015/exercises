package dynamic_programming

import "testing"

func TestUniquePathWithObstacles(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	ans := uniquePathsWithObstacles(obstacleGrid)
	if ans != 2 {
		t.Errorf("uniquePathsWithObstacles(obstacleGrid) = %d; want 2", ans)
	}
}
