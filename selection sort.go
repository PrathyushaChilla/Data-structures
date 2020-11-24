package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func main() {
 
    slice := generateSlice(20)
    fmt.Println("\n--- Unsorted --- \n\n", slice)
    selectionsort(slice)
    fmt.Println("\n--- Sorted ---", slice)
}
 
// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
 
    slice := make([]int, size, size)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(999) - rand.Intn(999)
    }
    return slice
}
  
func selectionsort(items []int) {
    var n = len(items)
    for i := 0; i < n; i++ {
        var minIdx = i
        for j := i; j < n; j++ {
            if items[j] < items[minIdx] {
                minIdx = j
            }
        }
        items[i], items[minIdx] = items[minIdx], items[i]
    }
}
output:
--- Unsorted --- 

 [-86 -816 -205 86 -106 -624 357 189 -263 149 419 125 -277 -72 235 117 -361 -321 528 -170]

--- Sorted --- [-816 -624 -361 -321 -277 -263 -205 -170 -106 -86 -72 86 117 125 149 189 235 357 419 528]

Program exited.