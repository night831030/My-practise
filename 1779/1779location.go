package location

import "math"

// import "fmt"

func main() {
	//     var point [][]int
	//     point[0][0] = 3
	//     point[0][1] = 5
	//     point[1][0] = 5
	//     point[0][1] = 7
	//     point[2][0] = 1000
	//     point[0][1] = 10
	//     point[3][0] = 50
	//     point[0][1] = -300
	//     point[4][0] = 5
	//     point[0][1] = 10
	point := [][]int{{3, 5}, {5, 7}, {1000, 10}, {5, -300}, {5, 10}}
	nearestValidPoint(5, 10, point)
}

func nearestValidPoint(x int, y int, points [][]int) int {
	z := -1                            ////預設為無同一線上座標，回傳-1
	ans := math.MaxInt64               ////最近距離預設int最大值
	for i := 0; i < len(points); i++ { //////以迴圈比對X,Y軸座標
		if points[i][0] == x || points[i][1] == y {
			a := abs(points[i][0]-x) + abs(points[i][1]-y) /////a為計算最小距離
			if a < ans {                                   ////最小距離為答案，如小於現有最小值則取代
				ans = a
				z = i
			}
		}
	}
	return z

}

func abs(x int) int {
	if x < 0 { /////距離無負數
		return -x
	} else {
		return x
	}
}
