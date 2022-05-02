package main

func createfib() func(int) []int {
	fList := []int{0, 1, 1, 2, 3, 5}
	return func(n int) []int {
		if n > len(fList) {
			for n > len(fList) {
				lastIndex := len(fList) - 1
				fList = append(fList, fList[lastIndex]+fList[lastIndex-1])
			}
		}
		return fList[:n]
	}
}

func profileTime(f func(int) []int) func(int) []int {

	return func(a int) []int {
		result := f(a)
		return result
	}
}
func main() {
	fib := createFib()
	fib = profileTime(fib)
	fib(60)
	fib(600)
}
