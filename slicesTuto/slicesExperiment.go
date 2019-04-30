package sliceexplanation

// PrintMySlice is just a function a la con
func PrintMySlice() []int {
	var numList []int
	numList = append(numList, 1, 2, 3, 4, 5, 6) // This is how you append a slice(I don't like it!)
	return numList
}
