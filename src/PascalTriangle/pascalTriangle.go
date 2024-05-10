package main

func main() {
	generate(5)
}

func generate(numRows int) [][]int {
	var arr [][]int

	for i := 0; i < numRows; i++ {
		var arIn []int
		arIn = append(arIn, 1)
		for j := 1; j <= i; j++ {
			if i == 1 {
				arIn = append(arIn, arr[i-1][j-1])
				continue
			}
			if j == i {
				arIn = append(arIn, 1)
				continue
			}
			arIn = append(arIn, arr[i-1][j-1]+arr[i-1][j])
		}
		arr = append(arr, arIn)
	}
	return arr
}
