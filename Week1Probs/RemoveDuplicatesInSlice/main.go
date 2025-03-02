package main

import "fmt"

func RemoveDUplicate(input []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, val := range input {
		if _, ok := seen[val]; !ok {        //logc to remember
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}

func main() {
	input := []int{3, 2, 2, 3, 4, 5,6,7,6,5,5,5,5,6,7}
	ans := RemoveDUplicate(input)
	fmt.Println("Slice after removing duplicates is : ",ans)
}


//CONCEPT OF PREALLOCATION,BELOW CODE TO JUDGE PERFORMANCE(SEE ON THE GO)

// import "testing"

// func BenchmarkNoPreallocate(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         // Don't preallocate our initial slice
//         init := []int64{}
//         init = append(init, 5)
//     }
// }

// func BenchmarkPreallocate(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         // Preallocate our initial slice
//         init := make([]int64, 0, 1)
//         init = append(init, 5)
//     }
// }