package main

func generateTreeValues(upTo int) []int {
	if upTo == 8 {
		return []int{4, 2, 6, 1, 3, 5, 7, 8}
	}
	rv := make([]int, 0, upTo-1)
	for i := 1; i <= upTo; i++ {
		rv = append(rv, i)
	}

	return rv
}

func main() {
	tree := BuildTree(generateTreeValues(8))

	tree.Print()
}
